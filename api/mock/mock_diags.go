// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStorageDiagsServer,OpenStorageDiagsClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	grpc "google.golang.org/grpc"
)

// MockOpenStorageDiagsServer is a mock of OpenStorageDiagsServer interface.
type MockOpenStorageDiagsServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageDiagsServerMockRecorder
}

// MockOpenStorageDiagsServerMockRecorder is the mock recorder for MockOpenStorageDiagsServer.
type MockOpenStorageDiagsServerMockRecorder struct {
	mock *MockOpenStorageDiagsServer
}

// NewMockOpenStorageDiagsServer creates a new mock instance.
func NewMockOpenStorageDiagsServer(ctrl *gomock.Controller) *MockOpenStorageDiagsServer {
	mock := &MockOpenStorageDiagsServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageDiagsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStorageDiagsServer) EXPECT() *MockOpenStorageDiagsServerMockRecorder {
	return m.recorder
}

// Collect mocks base method.
func (m *MockOpenStorageDiagsServer) Collect(arg0 context.Context, arg1 *api.SdkDiagsCollectRequest) (*api.SdkDiagsCollectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collect", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkDiagsCollectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect.
func (mr *MockOpenStorageDiagsServerMockRecorder) Collect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockOpenStorageDiagsServer)(nil).Collect), arg0, arg1)
}

// MockOpenStorageDiagsClient is a mock of OpenStorageDiagsClient interface.
type MockOpenStorageDiagsClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageDiagsClientMockRecorder
}

// MockOpenStorageDiagsClientMockRecorder is the mock recorder for MockOpenStorageDiagsClient.
type MockOpenStorageDiagsClientMockRecorder struct {
	mock *MockOpenStorageDiagsClient
}

// NewMockOpenStorageDiagsClient creates a new mock instance.
func NewMockOpenStorageDiagsClient(ctrl *gomock.Controller) *MockOpenStorageDiagsClient {
	mock := &MockOpenStorageDiagsClient{ctrl: ctrl}
	mock.recorder = &MockOpenStorageDiagsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStorageDiagsClient) EXPECT() *MockOpenStorageDiagsClientMockRecorder {
	return m.recorder
}

// Collect mocks base method.
func (m *MockOpenStorageDiagsClient) Collect(arg0 context.Context, arg1 *api.SdkDiagsCollectRequest, arg2 ...grpc.CallOption) (*api.SdkDiagsCollectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Collect", varargs...)
	ret0, _ := ret[0].(*api.SdkDiagsCollectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect.
func (mr *MockOpenStorageDiagsClientMockRecorder) Collect(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockOpenStorageDiagsClient)(nil).Collect), varargs...)
}