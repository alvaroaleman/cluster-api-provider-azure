/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../vmssextensions.go

// Package mock_vmssextensions is a generated GoMock package.
package mock_vmssextensions

import (
	context "context"
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
)

// MockVMSSExtensionScope is a mock of VMSSExtensionScope interface.
type MockVMSSExtensionScope struct {
	ctrl     *gomock.Controller
	recorder *MockVMSSExtensionScopeMockRecorder
}

// MockVMSSExtensionScopeMockRecorder is the mock recorder for MockVMSSExtensionScope.
type MockVMSSExtensionScopeMockRecorder struct {
	mock *MockVMSSExtensionScope
}

// NewMockVMSSExtensionScope creates a new mock instance.
func NewMockVMSSExtensionScope(ctrl *gomock.Controller) *MockVMSSExtensionScope {
	mock := &MockVMSSExtensionScope{ctrl: ctrl}
	mock.recorder = &MockVMSSExtensionScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVMSSExtensionScope) EXPECT() *MockVMSSExtensionScopeMockRecorder {
	return m.recorder
}

// AdditionalTags mocks base method.
func (m *MockVMSSExtensionScope) AdditionalTags() v1beta1.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1beta1.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockVMSSExtensionScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockVMSSExtensionScope)(nil).AdditionalTags))
}

// Authorizer mocks base method.
func (m *MockVMSSExtensionScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockVMSSExtensionScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockVMSSExtensionScope)(nil).Authorizer))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockVMSSExtensionScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockVMSSExtensionScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockVMSSExtensionScope)(nil).AvailabilitySetEnabled))
}

// BaseURI mocks base method.
func (m *MockVMSSExtensionScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockVMSSExtensionScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockVMSSExtensionScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockVMSSExtensionScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockVMSSExtensionScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockVMSSExtensionScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockVMSSExtensionScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockVMSSExtensionScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockVMSSExtensionScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockVMSSExtensionScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockVMSSExtensionScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockVMSSExtensionScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockVMSSExtensionScope) CloudProviderConfigOverrides() *v1beta1.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1beta1.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockVMSSExtensionScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockVMSSExtensionScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockVMSSExtensionScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockVMSSExtensionScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockVMSSExtensionScope)(nil).ClusterName))
}

// FailureDomains mocks base method.
func (m *MockVMSSExtensionScope) FailureDomains() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].([]string)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockVMSSExtensionScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockVMSSExtensionScope)(nil).FailureDomains))
}

// HashKey mocks base method.
func (m *MockVMSSExtensionScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockVMSSExtensionScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockVMSSExtensionScope)(nil).HashKey))
}

// Location mocks base method.
func (m *MockVMSSExtensionScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockVMSSExtensionScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockVMSSExtensionScope)(nil).Location))
}

// ResourceGroup mocks base method.
func (m *MockVMSSExtensionScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockVMSSExtensionScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockVMSSExtensionScope)(nil).ResourceGroup))
}

// SetBootstrapConditions mocks base method.
func (m *MockVMSSExtensionScope) SetBootstrapConditions(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBootstrapConditions", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBootstrapConditions indicates an expected call of SetBootstrapConditions.
func (mr *MockVMSSExtensionScopeMockRecorder) SetBootstrapConditions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBootstrapConditions", reflect.TypeOf((*MockVMSSExtensionScope)(nil).SetBootstrapConditions), arg0, arg1, arg2)
}

// SubscriptionID mocks base method.
func (m *MockVMSSExtensionScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockVMSSExtensionScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockVMSSExtensionScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockVMSSExtensionScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockVMSSExtensionScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockVMSSExtensionScope)(nil).TenantID))
}

// VMSSExtensionSpecs mocks base method.
func (m *MockVMSSExtensionScope) VMSSExtensionSpecs() []azure.ExtensionSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMSSExtensionSpecs")
	ret0, _ := ret[0].([]azure.ExtensionSpec)
	return ret0
}

// VMSSExtensionSpecs indicates an expected call of VMSSExtensionSpecs.
func (mr *MockVMSSExtensionScopeMockRecorder) VMSSExtensionSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMSSExtensionSpecs", reflect.TypeOf((*MockVMSSExtensionScope)(nil).VMSSExtensionSpecs))
}
