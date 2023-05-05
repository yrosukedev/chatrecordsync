// Code generated by MockGen. DO NOT EDIT.
// Source: chat_sync/wecom/transformer/field_transformer.go

// Package transformer is a generated GoMock package.
package transformer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	business "github.com/yrosukedev/chat_record_sync/chat_sync/business"
	wecom "github.com/yrosukedev/chat_record_sync/chat_sync/wecom"
)

// MockFieldTransformer is a mock of FieldTransformer interface.
type MockFieldTransformer struct {
	ctrl     *gomock.Controller
	recorder *MockFieldTransformerMockRecorder
}

// MockFieldTransformerMockRecorder is the mock recorder for MockFieldTransformer.
type MockFieldTransformerMockRecorder struct {
	mock *MockFieldTransformer
}

// NewMockFieldTransformer creates a new mock instance.
func NewMockFieldTransformer(ctrl *gomock.Controller) *MockFieldTransformer {
	mock := &MockFieldTransformer{ctrl: ctrl}
	mock.recorder = &MockFieldTransformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFieldTransformer) EXPECT() *MockFieldTransformerMockRecorder {
	return m.recorder
}

// Transform mocks base method.
func (m *MockFieldTransformer) Transform(wecomRecord *wecom.ChatRecord, chatRecord *business.ChatRecord) (*business.ChatRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transform", wecomRecord, chatRecord)
	ret0, _ := ret[0].(*business.ChatRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transform indicates an expected call of Transform.
func (mr *MockFieldTransformerMockRecorder) Transform(wecomRecord, chatRecord interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transform", reflect.TypeOf((*MockFieldTransformer)(nil).Transform), wecomRecord, chatRecord)
}
