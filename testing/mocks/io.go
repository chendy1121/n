// Code generated by MockGen. DO NOT EDIT.
// Source: io (interfaces: Reader,Writer)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Reader is a mock of Reader interface
type Reader struct {
	ctrl     *gomock.Controller
	recorder *ReaderMockRecorder
}

// ReaderMockRecorder is the mock recorder for Reader
type ReaderMockRecorder struct {
	mock *Reader
}

// NewReader creates a new mock instance
func NewReader(ctrl *gomock.Controller) *Reader {
	mock := &Reader{ctrl: ctrl}
	mock.recorder = &ReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Reader) EXPECT() *ReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *Reader) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *ReaderMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*Reader)(nil).Read), arg0)
}

// Writer is a mock of Writer interface
type Writer struct {
	ctrl     *gomock.Controller
	recorder *WriterMockRecorder
}

// WriterMockRecorder is the mock recorder for Writer
type WriterMockRecorder struct {
	mock *Writer
}

// NewWriter creates a new mock instance
func NewWriter(ctrl *gomock.Controller) *Writer {
	mock := &Writer{ctrl: ctrl}
	mock.recorder = &WriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Writer) EXPECT() *WriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *Writer) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *WriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*Writer)(nil).Write), arg0)
}
