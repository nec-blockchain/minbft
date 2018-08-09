// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger-labs/minbft/core/internal/messagelog (interfaces: MessageLog)

// Package mock_messagelog is a generated GoMock package.
package mock_messagelog

import (
	gomock "github.com/golang/mock/gomock"
	messages "github.com/hyperledger-labs/minbft/messages"
	reflect "reflect"
)

// MockMessageLog is a mock of MessageLog interface
type MockMessageLog struct {
	ctrl     *gomock.Controller
	recorder *MockMessageLogMockRecorder
}

// MockMessageLogMockRecorder is the mock recorder for MockMessageLog
type MockMessageLogMockRecorder struct {
	mock *MockMessageLog
}

// NewMockMessageLog creates a new mock instance
func NewMockMessageLog(ctrl *gomock.Controller) *MockMessageLog {
	mock := &MockMessageLog{ctrl: ctrl}
	mock.recorder = &MockMessageLogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageLog) EXPECT() *MockMessageLogMockRecorder {
	return m.recorder
}

// Append mocks base method
func (m *MockMessageLog) Append(arg0 *messages.Message) {
	m.ctrl.Call(m, "Append", arg0)
}

// Append indicates an expected call of Append
func (mr *MockMessageLogMockRecorder) Append(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockMessageLog)(nil).Append), arg0)
}

// Stream mocks base method
func (m *MockMessageLog) Stream(arg0 <-chan struct{}) <-chan *messages.Message {
	ret := m.ctrl.Call(m, "Stream", arg0)
	ret0, _ := ret[0].(<-chan *messages.Message)
	return ret0
}

// Stream indicates an expected call of Stream
func (mr *MockMessageLogMockRecorder) Stream(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockMessageLog)(nil).Stream), arg0)
}
