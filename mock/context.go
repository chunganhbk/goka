// Automatically generated by MockGen. DO NOT EDIT!
// Source: stash.lvint.de/lab/goka (interfaces: Context)

package mock

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Context interface
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *_MockContextRecorder
}

// Recorder for MockContext (not exported)
type _MockContextRecorder struct {
	mock *MockContext
}

func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &_MockContextRecorder{mock}
	return mock
}

func (_m *MockContext) EXPECT() *_MockContextRecorder {
	return _m.recorder
}

func (_m *MockContext) Emit(_param0 string, _param1 string, _param2 []byte) error {
	ret := _m.ctrl.Call(_m, "Emit", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockContextRecorder) Emit(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Emit", arg0, arg1, arg2)
}

func (_m *MockContext) Fail(_param0 error) {
	_m.ctrl.Call(_m, "Fail", _param0)
}

func (_mr *_MockContextRecorder) Fail(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fail", arg0)
}

func (_m *MockContext) Join(_param0 string) (interface{}, error) {
	ret := _m.ctrl.Call(_m, "Join", _param0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContextRecorder) Join(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Join", arg0)
}

func (_m *MockContext) Key() string {
	ret := _m.ctrl.Call(_m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockContextRecorder) Key() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Key")
}

func (_m *MockContext) Loopback(_param0 string, _param1 interface{}) error {
	ret := _m.ctrl.Call(_m, "Loopback", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockContextRecorder) Loopback(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Loopback", arg0, arg1)
}

func (_m *MockContext) SetValue(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "SetValue", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockContextRecorder) SetValue(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetValue", arg0)
}

func (_m *MockContext) Topic() string {
	ret := _m.ctrl.Call(_m, "Topic")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockContextRecorder) Topic() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Topic")
}

func (_m *MockContext) Value() (interface{}, error) {
	ret := _m.ctrl.Call(_m, "Value")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockContextRecorder) Value() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Value")
}
