// Code generated by MockGen. DO NOT EDIT.
// Source: ./sync.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/spacemeshos/go-spacemesh/common/types"
)

// MockSyncStateProvider is a mock of SyncStateProvider interface.
type MockSyncStateProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSyncStateProviderMockRecorder
}

// MockSyncStateProviderMockRecorder is the mock recorder for MockSyncStateProvider.
type MockSyncStateProviderMockRecorder struct {
	mock *MockSyncStateProvider
}

// NewMockSyncStateProvider creates a new mock instance.
func NewMockSyncStateProvider(ctrl *gomock.Controller) *MockSyncStateProvider {
	mock := &MockSyncStateProvider{ctrl: ctrl}
	mock.recorder = &MockSyncStateProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncStateProvider) EXPECT() *MockSyncStateProviderMockRecorder {
	return m.recorder
}

// IsBeaconSynced mocks base method.
func (m *MockSyncStateProvider) IsBeaconSynced(arg0 types.EpochID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBeaconSynced", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBeaconSynced indicates an expected call of IsBeaconSynced.
func (mr *MockSyncStateProviderMockRecorder) IsBeaconSynced(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBeaconSynced", reflect.TypeOf((*MockSyncStateProvider)(nil).IsBeaconSynced), arg0)
}

// IsSynced mocks base method.
func (m *MockSyncStateProvider) IsSynced(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSynced", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSynced indicates an expected call of IsSynced.
func (mr *MockSyncStateProviderMockRecorder) IsSynced(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSynced", reflect.TypeOf((*MockSyncStateProvider)(nil).IsSynced), arg0)
}

// SyncedBefore mocks base method.
func (m *MockSyncStateProvider) SyncedBefore(arg0 types.EpochID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncedBefore", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SyncedBefore indicates an expected call of SyncedBefore.
func (mr *MockSyncStateProviderMockRecorder) SyncedBefore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncedBefore", reflect.TypeOf((*MockSyncStateProvider)(nil).SyncedBefore), arg0)
}
