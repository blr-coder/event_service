// Code generated by MockGen. DO NOT EDIT.
// Source: event_service/internal/event/repositories (interfaces: IEventRepository)

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	repository_models "event_service/internal/event/repositories/repository_models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIEventRepository is a mock of IEventRepository interface.
type MockIEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEventRepositoryMockRecorder
}

// MockIEventRepositoryMockRecorder is the mock recorder for MockIEventRepository.
type MockIEventRepositoryMockRecorder struct {
	mock *MockIEventRepository
}

// NewMockIEventRepository creates a new mock instance.
func NewMockIEventRepository(ctrl *gomock.Controller) *MockIEventRepository {
	mock := &MockIEventRepository{ctrl: ctrl}
	mock.recorder = &MockIEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEventRepository) EXPECT() *MockIEventRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIEventRepository) Create(arg0 context.Context, arg1 *repository_models.CreateEventRepositoryDTO) (*repository_models.EventRepositoryDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*repository_models.EventRepositoryDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIEventRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIEventRepository)(nil).Create), arg0, arg1)
}

// List mocks base method.
func (m *MockIEventRepository) List(arg0 context.Context, arg1 *repository_models.EventRepositoryFilter) ([]*repository_models.EventRepositoryDTO, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*repository_models.EventRepositoryDTO)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockIEventRepositoryMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIEventRepository)(nil).List), arg0, arg1)
}
