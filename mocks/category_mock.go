// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/category.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/category.go -destination=mocks/category_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	model "spun/internal/model"

	gomock "go.uber.org/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryRepository) Create(category *model.Category) (*model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", category)
	ret0, _ := ret[0].(*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCategoryRepositoryMockRecorder) Create(category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryRepository)(nil).Create), category)
}

// Delete mocks base method.
func (m *MockCategoryRepository) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategoryRepository)(nil).Delete), id)
}

// GetByID mocks base method.
func (m *MockCategoryRepository) GetByID(id int64) (*model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCategoryRepositoryMockRecorder) GetByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCategoryRepository)(nil).GetByID), id)
}

// List mocks base method.
func (m *MockCategoryRepository) List(filter map[string]any, pagination *model.Pagination) ([]*model.Category, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", filter, pagination)
	ret0, _ := ret[0].([]*model.Category)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockCategoryRepositoryMockRecorder) List(filter, pagination any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCategoryRepository)(nil).List), filter, pagination)
}

// Update mocks base method.
func (m *MockCategoryRepository) Update(id int64, original, modified *model.Category) (*model.Category, map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, original, modified)
	ret0, _ := ret[0].(*model.Category)
	ret1, _ := ret[1].(map[string]any)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockCategoryRepositoryMockRecorder) Update(id, original, modified any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryRepository)(nil).Update), id, original, modified)
}