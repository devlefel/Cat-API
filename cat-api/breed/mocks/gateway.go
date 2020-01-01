// Code generated by MockGen. DO NOT EDIT.
// Source: cat-api/breed/gateway.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "cat-api/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGateway is a mock of Gateway interface
type MockGateway struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayMockRecorder
}

// MockGatewayMockRecorder is the mock recorder for MockGateway
type MockGatewayMockRecorder struct {
	mock *MockGateway
}

// NewMockGateway creates a new mock instance
func NewMockGateway(ctrl *gomock.Controller) *MockGateway {
	mock := &MockGateway{ctrl: ctrl}
	mock.recorder = &MockGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGateway) EXPECT() *MockGatewayMockRecorder {
	return m.recorder
}

// GetBreedByName mocks base method
func (m *MockGateway) GetBreedByName(breedName string) (*models.Breed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBreedByName", breedName)
	ret0, _ := ret[0].(*models.Breed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBreedByName indicates an expected call of GetBreedByName
func (mr *MockGatewayMockRecorder) GetBreedByName(breedName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBreedByName", reflect.TypeOf((*MockGateway)(nil).GetBreedByName), breedName)
}
