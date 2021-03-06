// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bagusays/efishery-technical-test/internal/service (interfaces: Resource)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/bagusays/efishery-technical-test/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockResource is a mock of Resource interface.
type MockResource struct {
	ctrl     *gomock.Controller
	recorder *MockResourceMockRecorder
}

// MockResourceMockRecorder is the mock recorder for MockResource.
type MockResourceMockRecorder struct {
	mock *MockResource
}

// NewMockResource creates a new mock instance.
func NewMockResource(ctrl *gomock.Controller) *MockResource {
	mock := &MockResource{ctrl: ctrl}
	mock.recorder = &MockResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResource) EXPECT() *MockResourceMockRecorder {
	return m.recorder
}

// FetchResource mocks base method.
func (m *MockResource) FetchResource(arg0 context.Context) ([]model.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchResource", arg0)
	ret0, _ := ret[0].([]model.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchResource indicates an expected call of FetchResource.
func (mr *MockResourceMockRecorder) FetchResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchResource", reflect.TypeOf((*MockResource)(nil).FetchResource), arg0)
}

// ResourceStatistics mocks base method.
func (m *MockResource) ResourceStatistics(arg0 context.Context) ([]model.ResourceStatistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceStatistics", arg0)
	ret0, _ := ret[0].([]model.ResourceStatistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourceStatistics indicates an expected call of ResourceStatistics.
func (mr *MockResourceMockRecorder) ResourceStatistics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceStatistics", reflect.TypeOf((*MockResource)(nil).ResourceStatistics), arg0)
}
