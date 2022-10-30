// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/kirill0909/resource-keeper-api/models"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), email, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// MockUserResource is a mock of UserResource interface.
type MockUserResource struct {
	ctrl     *gomock.Controller
	recorder *MockUserResourceMockRecorder
}

// MockUserResourceMockRecorder is the mock recorder for MockUserResource.
type MockUserResourceMockRecorder struct {
	mock *MockUserResource
}

// NewMockUserResource creates a new mock instance.
func NewMockUserResource(ctrl *gomock.Controller) *MockUserResource {
	mock := &MockUserResource{ctrl: ctrl}
	mock.recorder = &MockUserResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserResource) EXPECT() *MockUserResourceMockRecorder {
	return m.recorder
}

// CreateResource mocks base method.
func (m *MockUserResource) CreateResource(resource models.UserResource) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResource", resource)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateResource indicates an expected call of CreateResource.
func (mr *MockUserResourceMockRecorder) CreateResource(resource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResource", reflect.TypeOf((*MockUserResource)(nil).CreateResource), resource)
}

// DeleteResource mocks base method.
func (m *MockUserResource) DeleteResource(userId, resourceId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResource", userId, resourceId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteResource indicates an expected call of DeleteResource.
func (mr *MockUserResourceMockRecorder) DeleteResource(userId, resourceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResource", reflect.TypeOf((*MockUserResource)(nil).DeleteResource), userId, resourceId)
}

// GetAllResources mocks base method.
func (m *MockUserResource) GetAllResources(userId int) ([]models.UserResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllResources", userId)
	ret0, _ := ret[0].([]models.UserResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllResources indicates an expected call of GetAllResources.
func (mr *MockUserResourceMockRecorder) GetAllResources(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllResources", reflect.TypeOf((*MockUserResource)(nil).GetAllResources), userId)
}

// GetById mocks base method.
func (m *MockUserResource) GetById(userId, resourceId int) (models.UserResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, resourceId)
	ret0, _ := ret[0].(models.UserResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUserResourceMockRecorder) GetById(userId, resourceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserResource)(nil).GetById), userId, resourceId)
}

// UpdateResource mocks base method.
func (m *MockUserResource) UpdateResource(userId, resourceId int, input models.UserResourceUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResource", userId, resourceId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResource indicates an expected call of UpdateResource.
func (mr *MockUserResourceMockRecorder) UpdateResource(userId, resourceId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResource", reflect.TypeOf((*MockUserResource)(nil).UpdateResource), userId, resourceId, input)
}
