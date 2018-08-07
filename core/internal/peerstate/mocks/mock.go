// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nec-blockchain/minbft/core/internal/peerstate (interfaces: State)

// Package mock_peerstate is a generated GoMock package.
package mock_peerstate

import (
	gomock "github.com/golang/mock/gomock"
	usig "github.com/nec-blockchain/minbft/usig"
	reflect "reflect"
)

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// CaptureUI mocks base method
func (m *MockState) CaptureUI(arg0 *usig.UI) bool {
	ret := m.ctrl.Call(m, "CaptureUI", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CaptureUI indicates an expected call of CaptureUI
func (mr *MockStateMockRecorder) CaptureUI(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureUI", reflect.TypeOf((*MockState)(nil).CaptureUI), arg0)
}

// ReleaseUI mocks base method
func (m *MockState) ReleaseUI(arg0 *usig.UI) error {
	ret := m.ctrl.Call(m, "ReleaseUI", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseUI indicates an expected call of ReleaseUI
func (mr *MockStateMockRecorder) ReleaseUI(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseUI", reflect.TypeOf((*MockState)(nil).ReleaseUI), arg0)
}
