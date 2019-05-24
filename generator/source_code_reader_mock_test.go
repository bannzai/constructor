// Code generated by MockGen. DO NOT EDIT.
// Source: generator/source_code_reader.go

// Package generator is a generated GoMock package.
package generator

import (
	structure "github.com/constructor/structure"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// SourceCodeReaderMock is a mock of SourceCodeReader interface
type SourceCodeReaderMock struct {
	ctrl     *gomock.Controller
	recorder *SourceCodeReaderMockMockRecorder
}

// SourceCodeReaderMockMockRecorder is the mock recorder for SourceCodeReaderMock
type SourceCodeReaderMockMockRecorder struct {
	mock *SourceCodeReaderMock
}

// NewSourceCodeReaderMock creates a new mock instance
func NewSourceCodeReaderMock(ctrl *gomock.Controller) *SourceCodeReaderMock {
	mock := &SourceCodeReaderMock{ctrl: ctrl}
	mock.recorder = &SourceCodeReaderMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *SourceCodeReaderMock) EXPECT() *SourceCodeReaderMockMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *SourceCodeReaderMock) Read(filePath structure.Path) structure.Code {
	ret := m.ctrl.Call(m, "Read", filePath)
	ret0, _ := ret[0].(structure.Code)
	return ret0
}

// Read indicates an expected call of Read
func (mr *SourceCodeReaderMockMockRecorder) Read(filePath interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*SourceCodeReaderMock)(nil).Read), filePath)
}
