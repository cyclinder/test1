// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: ./lease_election.go

// Package mock_election is a generated GoMock package.
package mock_election

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kubernetes "k8s.io/client-go/kubernetes"
)

// MockSpiderLeaseElector is a mock of SpiderLeaseElector interface.
type MockSpiderLeaseElector struct {
	ctrl     *gomock.Controller
	recorder *MockSpiderLeaseElectorMockRecorder
}

// MockSpiderLeaseElectorMockRecorder is the mock recorder for MockSpiderLeaseElector.
type MockSpiderLeaseElectorMockRecorder struct {
	mock *MockSpiderLeaseElector
}

// NewMockSpiderLeaseElector creates a new mock instance.
func NewMockSpiderLeaseElector(ctrl *gomock.Controller) *MockSpiderLeaseElector {
	mock := &MockSpiderLeaseElector{ctrl: ctrl}
	mock.recorder = &MockSpiderLeaseElectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpiderLeaseElector) EXPECT() *MockSpiderLeaseElectorMockRecorder {
	return m.recorder
}

// GetLeader mocks base method.
func (m *MockSpiderLeaseElector) GetLeader() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeader")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLeader indicates an expected call of GetLeader.
func (mr *MockSpiderLeaseElectorMockRecorder) GetLeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeader", reflect.TypeOf((*MockSpiderLeaseElector)(nil).GetLeader))
}

// IsElected mocks base method.
func (m *MockSpiderLeaseElector) IsElected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsElected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsElected indicates an expected call of IsElected.
func (mr *MockSpiderLeaseElectorMockRecorder) IsElected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsElected", reflect.TypeOf((*MockSpiderLeaseElector)(nil).IsElected))
}

// Run mocks base method.
func (m *MockSpiderLeaseElector) Run(ctx context.Context, clientSet kubernetes.Interface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, clientSet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockSpiderLeaseElectorMockRecorder) Run(ctx, clientSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockSpiderLeaseElector)(nil).Run), ctx, clientSet)
}
