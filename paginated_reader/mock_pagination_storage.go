// Code generated by MockGen. DO NOT EDIT.
// Source: paginated_reader/pagination_storage.go

// Package paginated_reader is a generated GoMock package.
package paginated_reader

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChatRecordPaginationStorage is a mock of ChatRecordPaginationStorage interface.
type MockChatRecordPaginationStorage struct {
	ctrl     *gomock.Controller
	recorder *MockChatRecordPaginationStorageMockRecorder
}

// MockChatRecordPaginationStorageMockRecorder is the mock recorder for MockChatRecordPaginationStorage.
type MockChatRecordPaginationStorageMockRecorder struct {
	mock *MockChatRecordPaginationStorage
}

// NewMockChatRecordPaginationStorage creates a new mock instance.
func NewMockChatRecordPaginationStorage(ctrl *gomock.Controller) *MockChatRecordPaginationStorage {
	mock := &MockChatRecordPaginationStorage{ctrl: ctrl}
	mock.recorder = &MockChatRecordPaginationStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatRecordPaginationStorage) EXPECT() *MockChatRecordPaginationStorageMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockChatRecordPaginationStorage) Get() (PageToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(PageToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockChatRecordPaginationStorageMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockChatRecordPaginationStorage)(nil).Get))
}

// Set mocks base method.
func (m *MockChatRecordPaginationStorage) Set(pageToken PageToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", pageToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockChatRecordPaginationStorageMockRecorder) Set(pageToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockChatRecordPaginationStorage)(nil).Set), pageToken)
}
