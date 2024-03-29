// Code generated by MockGen. DO NOT EDIT.
// Source: cat-api/breed/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "cat-api/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetBreedByName mocks base method
func (m *MockRepository) GetBreedByName(breedName string) (*models.Breed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBreedByName", breedName)
	ret0, _ := ret[0].(*models.Breed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBreedByName indicates an expected call of GetBreedByName
func (mr *MockRepositoryMockRecorder) GetBreedByName(breedName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBreedByName", reflect.TypeOf((*MockRepository)(nil).GetBreedByName), breedName)
}
