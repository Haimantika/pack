// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/pack/internal/dist (interfaces: Buildpack)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	dist "github.com/buildpacks/pack/internal/dist"
)

// MockBuildpack is a mock of Buildpack interface.
type MockBuildpack struct {
	ctrl     *gomock.Controller
	recorder *MockBuildpackMockRecorder
}

// MockBuildpackMockRecorder is the mock recorder for MockBuildpack.
type MockBuildpackMockRecorder struct {
	mock *MockBuildpack
}

// NewMockBuildpack creates a new mock instance.
func NewMockBuildpack(ctrl *gomock.Controller) *MockBuildpack {
	mock := &MockBuildpack{ctrl: ctrl}
	mock.recorder = &MockBuildpackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildpack) EXPECT() *MockBuildpackMockRecorder {
	return m.recorder
}

// Descriptor mocks base method.
func (m *MockBuildpack) Descriptor() dist.BuildpackDescriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].(dist.BuildpackDescriptor)
	return ret0
}

// Descriptor indicates an expected call of Descriptor.
func (mr *MockBuildpackMockRecorder) Descriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*MockBuildpack)(nil).Descriptor))
}

// Open mocks base method.
func (m *MockBuildpack) Open() (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockBuildpackMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockBuildpack)(nil).Open))
}
