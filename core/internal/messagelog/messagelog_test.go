// Copyright (c) 2018 NEC Laboratories Europe GmbH.
//
// Authors: Sergey Fedorov <sergey.fedorov@neclab.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package messagelog

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperledger-labs/minbft/messages"
)

func TestAppend(t *testing.T) {
	log := New()

	done := make(chan struct{})
	defer close(done)

	log.Append(makeMsg())

	// Should not block if there is no stream
	log.Append(makeMsg())

	_ = log.Stream(done)

	// Should not block if there is a stream
	log.Append(makeMsg())
	log.Append(makeMsg())

	_ = log.Stream(done)

	// Should not block if there are multiple streams
	log.Append(makeMsg())
	log.Append(makeMsg())
}

func TestStream(t *testing.T) {
	const nrMessages = 5

	log := New()
	msgs := makeManyMsgs(nrMessages)

	done := make(chan struct{})
	ch1 := log.Stream(done)
	ch2 := log.Stream(done)

	for _, msg := range msgs {
		log.Append(msg)
	}

	receiveMessages := func(ch <-chan messages.Message) {
		for i, msg := range msgs {
			assert.Equalf(t, msg, <-ch, "Unexpected message %d", i)
		}
	}

	receiveMessages(ch1)
	receiveMessages(ch2)

	close(done)
	_, more := <-ch1
	assert.False(t, more, "Channel not closed")
	_, more = <-ch2
	assert.False(t, more, "Channel not closed")
}

func TestMessages(t *testing.T) {
	const nrMessages = 5

	log := New()
	msgs := makeManyMsgs(nrMessages)

	for _, msg := range msgs {
		log.Append(msg)
	}
	assert.Equalf(t, msgs, log.Messages(), "Unexpected messages")
}

func TestReset(t *testing.T) {
	const nrMessages = 23

	log := New()
	ch := log.Stream(nil)

	msgs := makeManyMsgs(nrMessages)
	log.Reset(msgs)

	for i, m := range msgs {
		assert.Equalf(t, m, <-ch, "Unexpected message %d", i)
	}

	msgs2 := makeManyMsgs(nrMessages)
	log.Reset(msgs2)

	ch2 := log.Stream(nil)
	for i, m := range msgs2 {
		assert.Equalf(t, m, <-ch, "Unexpected message %d", i)
		assert.Equalf(t, m, <-ch2, "Unexpected message %d", i)
	}
	assert.Equalf(t, msgs2, log.Messages(), "Unexpected messages")
}

func TestResetConcurrent(t *testing.T) {
	const nrStreams = 3
	const nrMessages = 23

	log := New()

	msgs := makeManyMsgs(nrMessages)
	msgs2 := makeManyMsgs(nrMessages)
	log.Reset(msgs)

	wg := new(sync.WaitGroup)
	wg.Add(nrStreams)
	for id := 0; id < nrStreams; id++ {
		ch := log.Stream(nil)

		go func(streamID int) {
			defer wg.Done()

			var i, j int
			for j < len(msgs2) {
				ok := assert.Conditionf(t, func() bool {
					m := <-ch
					switch {
					case i < len(msgs) && m == msgs[i]:
						i++
						return assert.Zero(t, j)
					case m == msgs2[j]:
						j++
						return true
					default:
						return false
					}
				}, "Unexpected message from stream %d", streamID)
				if !ok {
					break
				}
			}
		}(id)
	}

	log.Reset(msgs2)
	wg.Wait()
}

func TestConcurrent(t *testing.T) {
	const nrStreams = 3
	const nrMessages = 5

	log := New()
	msgs := makeManyMsgs(nrMessages)

	wg := new(sync.WaitGroup)
	wg.Add(nrStreams)
	for id := 0; id < nrStreams; id++ {
		go func(streamID int) {
			defer wg.Done()

			done := make(chan struct{})
			ch := log.Stream(done)

			for i, msg := range msgs {
				assert.Equalf(t, msg, <-ch, "Unexpected message %d from stream %d", i, streamID)
				assert.Equalf(t, msgs[:i], log.Messages()[:i], "Unexpected messages in the log")
			}

			close(done)
			_, more := <-ch
			assert.Falsef(t, more, "Stream %d channel not closed", streamID)
		}(id)
	}

	for _, msg := range msgs {
		log.Append(msg)
	}

	wg.Wait()
}

func TestWithFaulty(t *testing.T) {
	const nrStreams = 5
	const nrFaulty = 2
	const nrMessages = 10

	log := New()
	msgs := makeManyMsgs(nrMessages)

	wg := new(sync.WaitGroup)
	wg.Add(nrStreams)
	for id := 0; id < nrStreams; id++ {
		go func(streamID int) {
			done := make(chan struct{})
			defer close(done)
			ch := log.Stream(done)

			if streamID < nrFaulty {
				wg.Done()
				wg.Wait()
				return
			}

			for i, msg := range msgs {
				assert.Equalf(t, msg, <-ch, "Unexpected message %d from stream %d", i, streamID)
				assert.Equalf(t, msgs[:i], log.Messages()[:i], "Unexpected messages in the log")
			}

			wg.Done()
		}(id)
	}

	for _, msg := range msgs {
		log.Append(msg)
	}

	wg.Wait()
}

func makeManyMsgs(nrMessages int) []messages.Message {
	msgs := make([]messages.Message, nrMessages)
	for i := 0; i < nrMessages; i++ {
		msgs[i] = makeMsg()
	}
	return msgs
}

func makeMsg() messages.Message {
	return struct {
		messages.ReplicaMessage
		i int
	}{i: rand.Int()}
}
