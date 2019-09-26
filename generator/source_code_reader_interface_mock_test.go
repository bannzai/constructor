// Code generated by MockGen. DO NOT EDIT.
// Source: source_code_reader.go

// Package generator is a generated GoMock package.
package generator

import (
	structure "github.com/bannzai/constructor/structure"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSourceCodeReader is a mock of SourceCodeReader interface
type MockSourceCodeReader struct {
	ctrl     *gomock.Controller
	recorder *MockSourceCodeReaderMockRecorder
}

// MockSourceCodeReaderMockRecorder is the mock recorder for MockSourceCodeReader
type MockSourceCodeReaderMockRecorder struct {
	mock *MockSourceCodeReader
}

// NewMockSourceCodeReader creates a new mock instance
func NewMockSourceCodeReader(ctrl *gomock.Controller) *MockSourceCodeReader {
	mock := &MockSourceCodeReader{ctrl: ctrl}
	mock.recorder = &MockSourceCodeReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceCodeReader) EXPECT() *MockSourceCodeReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockSourceCodeReader) Read(filePath structure.Path, ignoreFieldNames []string) structure.Code {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", filePath, ignoreFieldNames)
	ret0, _ := ret[0].(structure.Code)
	return ret0
}

// Read indicates an expected call of Read
func (mr *MockSourceCodeReaderMockRecorder) Read(filePath, ignoreFieldNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSourceCodeReader)(nil).Read), filePath, ignoreFieldNames)
}

// ReadWithType mocks base method
func (m *MockSourceCodeReader) ReadWithType(filePath structure.Path, generatedTypeName string, ignoreFieldNames []string) structure.Code {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithType", filePath, generatedTypeName, ignoreFieldNames)
	ret0, _ := ret[0].(structure.Code)
	return ret0
}

// ReadWithType indicates an expected call of ReadWithType
func (mr *MockSourceCodeReaderMockRecorder) ReadWithType(filePath, generatedTypeName, ignoreFieldNames interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithType", reflect.TypeOf((*MockSourceCodeReader)(nil).ReadWithType), filePath, generatedTypeName, ignoreFieldNames)
}