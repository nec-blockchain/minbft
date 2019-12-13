// Copyright (c) 2018-2019 NEC Laboratories Europe GmbH.
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

package clientstate

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	testifymock "github.com/stretchr/testify/mock"

	"github.com/hyperledger-labs/minbft/core/internal/timer"

	timermock "github.com/hyperledger-labs/minbft/core/internal/timer/mock"
)

func TestTimeout(t *testing.T) {
	t.Run("Start", testStartTimeout)
	t.Run("Stop", testStopTimeout)
}

func testStartTimeout(t *testing.T) {
	mock := new(testifymock.Mock)
	defer mock.AssertExpectations(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s, timerProvider, handleRequestTimeout, handlePrepareTimeout := setupTimeoutMock(mock, ctrl)

	// Start with disabled timeout
	mock.On("requestTimeout").Return(time.Duration(0)).Once()
	s.StartRequestTimer(handleRequestTimeout)

	mock.On("prepareTimeout").Return(time.Duration(0)).Once()
	s.StartPrepareTimer(handlePrepareTimeout)

	// Start with enabled timeout
	timeout := randTimeout()
	mockTimer := timermock.NewMockTimer(ctrl)
	timerProvider.EXPECT().AfterFunc(timeout, gomock.Any()).DoAndReturn(
		func(d time.Duration, f func()) timer.Timer {
			f()
			return mockTimer
		},
	).Times(2)
	mock.On("requestTimeout").Return(timeout).Once()
	mock.On("requestTimeoutHandler").Once()
	s.StartRequestTimer(handleRequestTimeout)
	mock.On("prepareTimeout").Return(timeout).Once()
	mock.On("prepareTimeoutHandler").Once()
	s.StartPrepareTimer(handlePrepareTimeout)

	// Restart timeout
	mockTimer.EXPECT().Stop().Times(2)
	timeout = randTimeout()
	mockTimer = timermock.NewMockTimer(ctrl)
	timerProvider.EXPECT().AfterFunc(timeout, gomock.Any()).DoAndReturn(
		func(d time.Duration, f func()) timer.Timer {
			f()
			return mockTimer
		},
	).Times(2)
	mock.On("requestTimeout").Return(timeout).Once()
	mock.On("requestTimeoutHandler").Once()
	s.StartRequestTimer(handleRequestTimeout)
	mock.On("prepareTimeout").Return(timeout).Once()
	mock.On("prepareTimeoutHandler").Once()
	s.StartPrepareTimer(handlePrepareTimeout)
}

func testStopTimeout(t *testing.T) {
	mock := new(testifymock.Mock)
	defer mock.AssertExpectations(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s, timerProvider, handleRequestTimeout, handlePrepareTimeout := setupTimeoutMock(mock, ctrl)

	timeout := randTimeout()
	mock.On("requestTimeout").Return(timeout)
	mock.On("prepareTimeout").Return(timeout)

	// Stop before started
	assert.NotPanics(t, func() {
		s.StopRequestTimer()
		s.StopPrepareTimer()
	})

	// Start and stop request timer
	mockTimer := timermock.NewMockTimer(ctrl)
	timerProvider.EXPECT().AfterFunc(timeout, gomock.Any()).Return(mockTimer).Times(2)
	s.StartRequestTimer(handleRequestTimeout)
	s.StartPrepareTimer(handlePrepareTimeout)
	mockTimer.EXPECT().Stop().Times(2)
	s.StopRequestTimer()
	s.StopPrepareTimer()

	// Stop again
	mockTimer.EXPECT().Stop().Times(2)
	s.StopRequestTimer()
	s.StopPrepareTimer()
}

func setupTimeoutMock(mock *testifymock.Mock, ctrl *gomock.Controller) (state State, timerProvider *timermock.MockProvider, handleRequestTimeout func(), handlePrepareTimeout func()) {
	handleRequestTimeout = func() {
		mock.MethodCalled("requestTimeoutHandler")
	}
	handlePrepareTimeout = func() {
		mock.MethodCalled("prepareTimeoutHandler")
	}
	requestTimeout := func() time.Duration {
		args := mock.MethodCalled("requestTimeout")
		return args.Get(0).(time.Duration)
	}
	prepareTimeout := func() time.Duration {
		args := mock.MethodCalled("prepareTimeout")
		return args.Get(0).(time.Duration)
	}
	timerProvider = timermock.NewMockProvider(ctrl)
	state = New(requestTimeout, prepareTimeout, WithTimerProvider(timerProvider))
	return state, timerProvider, handleRequestTimeout, handlePrepareTimeout
}

func randTimeout() time.Duration {
	return time.Duration(rand.Intn(math.MaxInt32-1) + 1) // positive nonzero
}
