// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger-labs/minbft/usig (interfaces: USIG)

// Package mock_usig is a generated GoMock package.
package mock_usig

import (
	gomock "github.com/golang/mock/gomock"
	usig "github.com/hyperledger-labs/minbft/usig"
	reflect "reflect"
)

// MockUSIG is a mock of USIG interface
type MockUSIG struct {
	ctrl     *gomock.Controller
	recorder *MockUSIGMockRecorder
}

// MockUSIGMockRecorder is the mock recorder for MockUSIG
type MockUSIGMockRecorder struct {
	mock *MockUSIG
}

// NewMockUSIG creates a new mock instance
func NewMockUSIG(ctrl *gomock.Controller) *MockUSIG {
	mock := &MockUSIG{ctrl: ctrl}
	mock.recorder = &MockUSIGMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUSIG) EXPECT() *MockUSIGMockRecorder {
	return m.recorder
}

// CreateUI mocks base method
func (m *MockUSIG) CreateUI(arg0 []byte) (*usig.UI, error) {
	ret := m.ctrl.Call(m, "CreateUI", arg0)
	ret0, _ := ret[0].(*usig.UI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUI indicates an expected call of CreateUI
func (mr *MockUSIGMockRecorder) CreateUI(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUI", reflect.TypeOf((*MockUSIG)(nil).CreateUI), arg0)
}

// ID mocks base method
func (m *MockUSIG) ID() []byte {
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockUSIGMockRecorder) ID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockUSIG)(nil).ID))
}

// VerifyUI mocks base method
func (m *MockUSIG) VerifyUI(arg0 []byte, arg1 *usig.UI, arg2 []byte) error {
	ret := m.ctrl.Call(m, "VerifyUI", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUI indicates an expected call of VerifyUI
func (mr *MockUSIGMockRecorder) VerifyUI(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUI", reflect.TypeOf((*MockUSIG)(nil).VerifyUI), arg0, arg1, arg2)
}
