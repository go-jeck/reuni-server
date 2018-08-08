// Code generated by MockGen. DO NOT EDIT.
// Source: organization/organization_processor.go

// Package organization is a generated GoMock package.
package organization

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockprocessor is a mock of processor interface
type Mockprocessor struct {
	ctrl     *gomock.Controller
	recorder *MockprocessorMockRecorder
}

// MockprocessorMockRecorder is the mock recorder for Mockprocessor
type MockprocessorMockRecorder struct {
	mock *Mockprocessor
}

// NewMockprocessor creates a new mock instance
func NewMockprocessor(ctrl *gomock.Controller) *Mockprocessor {
	mock := &Mockprocessor{ctrl: ctrl}
	mock.recorder = &MockprocessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockprocessor) EXPECT() *MockprocessorMockRecorder {
	return m.recorder
}

// createNewOrganizationProcessor mocks base method
func (m *Mockprocessor) createNewOrganizationProcessor(orginizationName string, userId int64) error {
	ret := m.ctrl.Call(m, "createNewOrganizationProcessor", orginizationName, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// createNewOrganizationProcessor indicates an expected call of createNewOrganizationProcessor
func (mr *MockprocessorMockRecorder) createNewOrganizationProcessor(orginizationName, userId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createNewOrganizationProcessor", reflect.TypeOf((*Mockprocessor)(nil).createNewOrganizationProcessor), orginizationName, userId)
}

// addUserProcessor mocks base method
func (m *Mockprocessor) addUserProcessor(member *Member) error {
	ret := m.ctrl.Call(m, "addUserProcessor", member)
	ret0, _ := ret[0].(error)
	return ret0
}

// addUserProcessor indicates an expected call of addUserProcessor
func (mr *MockprocessorMockRecorder) addUserProcessor(member interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addUserProcessor", reflect.TypeOf((*Mockprocessor)(nil).addUserProcessor), member)
}
