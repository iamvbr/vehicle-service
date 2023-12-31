// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package stores is a generated GoMock package.
package stores

import (
	context "context"
	reflect "reflect"

	models "github.com/iamvbr/vehicle-service/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockCarStorage is a mock of CarStorage interface.
type MockCarStorage struct {
	ctrl     *gomock.Controller
	recorder *MockCarStorageMockRecorder
}

// MockCarStorageMockRecorder is the mock recorder for MockCarStorage.
type MockCarStorageMockRecorder struct {
	mock *MockCarStorage
}

// NewMockCarStorage creates a new mock instance.
func NewMockCarStorage(ctrl *gomock.Controller) *MockCarStorage {
	mock := &MockCarStorage{ctrl: ctrl}
	mock.recorder = &MockCarStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarStorage) EXPECT() *MockCarStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockCarStorage) Create(ctx context.Context, m *models.Car) (*models.Car, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCarStorageMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCarStorage)(nil).Create), ctx, m)
}

// Exists mocks base method.
func (m *MockCarStorage) Exists(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockCarStorageMockRecorder) Exists(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockCarStorage)(nil).Exists), ctx, id)
}

// Find mocks base method.
func (m *MockCarStorage) Find(ctx context.Context, category string) ([]*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, category)
	ret0, _ := ret[0].([]*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCarStorageMockRecorder) Find(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCarStorage)(nil).Find), ctx, category)
}

// Get mocks base method.
func (m *MockCarStorage) Get(ctx context.Context, id string) (*models.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCarStorageMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCarStorage)(nil).Get), ctx, id)
}

// Update mocks base method.
func (m_2 *MockCarStorage) Update(ctx context.Context, id string, m *models.Car) (*models.Car, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, id, m)
	ret0, _ := ret[0].(*models.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCarStorageMockRecorder) Update(ctx, id, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCarStorage)(nil).Update), ctx, id, m)
}
