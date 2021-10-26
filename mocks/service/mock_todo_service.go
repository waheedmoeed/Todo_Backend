// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/abdulwaheed/todobackend/service (interfaces: TodoService)

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTodoService is a mock of TodoService interface
type MockTodoService struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceMockRecorder
}

// MockTodoServiceMockRecorder is the mock recorder for MockTodoService
type MockTodoServiceMockRecorder struct {
	mock *MockTodoService
}

// NewMockTodoService creates a new mock instance
func NewMockTodoService(ctrl *gomock.Controller) *MockTodoService {
	mock := &MockTodoService{ctrl: ctrl}
	mock.recorder = &MockTodoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTodoService) EXPECT() *MockTodoServiceMockRecorder {
	return m.recorder
}

// AddTask mocks base method
func (m *MockTodoService) AddTask(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask
func (mr *MockTodoServiceMockRecorder) AddTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockTodoService)(nil).AddTask), arg0)
}

// GetTask mocks base method
func (m *MockTodoService) GetTask(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask
func (mr *MockTodoServiceMockRecorder) GetTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockTodoService)(nil).GetTask), arg0)
}
