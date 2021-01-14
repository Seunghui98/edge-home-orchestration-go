/*******************************************************************************
 * Copyright 2020 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

// Code generated by MockGen. DO NOT EDIT.
// Source: ./executor.go

// Package mocks is a generated GoMock package.
package mocks

import (
	executor "github.com/lf-edge/edge-home-orchestration-go/src/controller/servicemgr/executor"
	notification "github.com/lf-edge/edge-home-orchestration-go/src/controller/servicemgr/notification"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	client "github.com/lf-edge/edge-home-orchestration-go/src/restinterface/client"
)

// MockServiceExecutor is a mock of ServiceExecutor interface
type MockServiceExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockServiceExecutorMockRecorder
}

// MockServiceExecutorMockRecorder is the mock recorder for MockServiceExecutor
type MockServiceExecutorMockRecorder struct {
	mock *MockServiceExecutor
}

// NewMockServiceExecutor creates a new mock instance
func NewMockServiceExecutor(ctrl *gomock.Controller) *MockServiceExecutor {
	mock := &MockServiceExecutor{ctrl: ctrl}
	mock.recorder = &MockServiceExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceExecutor) EXPECT() *MockServiceExecutorMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockServiceExecutor) Execute(arg0 executor.ServiceExecutionInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockServiceExecutorMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockServiceExecutor)(nil).Execute), arg0)
}

// SetNotiImpl mocks base method
func (m *MockServiceExecutor) SetNotiImpl(noti notification.Notification) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNotiImpl", noti)
}

// SetNotiImpl indicates an expected call of SetNotiImpl
func (mr *MockServiceExecutorMockRecorder) SetNotiImpl(noti interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNotiImpl", reflect.TypeOf((*MockServiceExecutor)(nil).SetNotiImpl), noti)
}

// SetClient mocks base method
func (m *MockServiceExecutor) SetClient(clientAPI client.Clienter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", clientAPI)
}

// SetClient indicates an expected call of SetClient
func (mr *MockServiceExecutorMockRecorder) SetClient(clientAPI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockServiceExecutor)(nil).SetClient), clientAPI)
}
