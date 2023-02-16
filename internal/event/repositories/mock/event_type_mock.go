// Code generated by MockGen. DO NOT EDIT.
// Source: event_service/internal/event/repositories (interfaces: IEventTypeRepository)

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	repository_models "event_service/internal/event/repositories/repository_models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIEventTypeRepository is a mock of IEventTypeRepository interface.
type MockIEventTypeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEventTypeRepositoryMockRecorder
}

// MockIEventTypeRepositoryMockRecorder is the mock recorder for MockIEventTypeRepository.
type MockIEventTypeRepositoryMockRecorder struct {
	mock *MockIEventTypeRepository
}

// NewMockIEventTypeRepository creates a new mock instance.
func NewMockIEventTypeRepository(ctrl *gomock.Controller) *MockIEventTypeRepository {
	mock := &MockIEventTypeRepository{ctrl: ctrl}
	mock.recorder = &MockIEventTypeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEventTypeRepository) EXPECT() *MockIEventTypeRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIEventTypeRepository) Create(arg0 context.Context, arg1 *repository_models.CreateEventTypeRepositoryDTO) (*repository_models.EventTypeRepositoryDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*repository_models.EventTypeRepositoryDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIEventTypeRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIEventTypeRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIEventTypeRepository) Delete(arg0 context.Context, arg1 *repository_models.DeleteEventTypeRepositoryDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIEventTypeRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIEventTypeRepository)(nil).Delete), arg0, arg1)
}

// List mocks base method.
func (m *MockIEventTypeRepository) List(arg0 context.Context, arg1 *repository_models.EventTypeRepositoryFilter) ([]*repository_models.EventTypeRepositoryDTO, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*repository_models.EventTypeRepositoryDTO)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockIEventTypeRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIEventTypeRepository)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockIEventTypeRepository) Update(arg0 context.Context, arg1 *repository_models.UpdateEventTypeRepositoryDTO) (*repository_models.EventTypeRepositoryDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*repository_models.EventTypeRepositoryDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIEventTypeRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIEventTypeRepository)(nil).Update), arg0, arg1)
}