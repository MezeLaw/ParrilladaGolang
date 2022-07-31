// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package handler is a generated GoMock package.
package handler

import (
	comida "Parrillada/models/comida"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockComidasService is a mock of ComidasService interface.
type MockComidasService struct {
	ctrl     *gomock.Controller
	recorder *MockComidasServiceMockRecorder
}

// MockComidasServiceMockRecorder is the mock recorder for MockComidasService.
type MockComidasServiceMockRecorder struct {
	mock *MockComidasService
}

// NewMockComidasService creates a new mock instance.
func NewMockComidasService(ctrl *gomock.Controller) *MockComidasService {
	mock := &MockComidasService{ctrl: ctrl}
	mock.recorder = &MockComidasServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComidasService) EXPECT() *MockComidasServiceMockRecorder {
	return m.recorder
}

// AgregarPlato mocks base method.
func (m *MockComidasService) AgregarPlato(plato comida.IComida) []comida.IComida {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgregarPlato", plato)
	ret0, _ := ret[0].([]comida.IComida)
	return ret0
}

// AgregarPlato indicates an expected call of AgregarPlato.
func (mr *MockComidasServiceMockRecorder) AgregarPlato(plato interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgregarPlato", reflect.TypeOf((*MockComidasService)(nil).AgregarPlato), plato)
}

// ConsultarPlatosDisponibles mocks base method.
func (m *MockComidasService) ConsultarPlatosDisponibles() []comida.IComida {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsultarPlatosDisponibles")
	ret0, _ := ret[0].([]comida.IComida)
	return ret0
}

// ConsultarPlatosDisponibles indicates an expected call of ConsultarPlatosDisponibles.
func (mr *MockComidasServiceMockRecorder) ConsultarPlatosDisponibles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsultarPlatosDisponibles", reflect.TypeOf((*MockComidasService)(nil).ConsultarPlatosDisponibles))
}

// EliminarPlato mocks base method.
func (m *MockComidasService) EliminarPlato(id string) []comida.IComida {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EliminarPlato", id)
	ret0, _ := ret[0].([]comida.IComida)
	return ret0
}

// EliminarPlato indicates an expected call of EliminarPlato.
func (mr *MockComidasServiceMockRecorder) EliminarPlato(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EliminarPlato", reflect.TypeOf((*MockComidasService)(nil).EliminarPlato), id)
}

// OfertaVegetariana mocks base method.
func (m *MockComidasService) OfertaVegetariana() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfertaVegetariana")
	ret0, _ := ret[0].(bool)
	return ret0
}

// OfertaVegetariana indicates an expected call of OfertaVegetariana.
func (mr *MockComidasServiceMockRecorder) OfertaVegetariana() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfertaVegetariana", reflect.TypeOf((*MockComidasService)(nil).OfertaVegetariana))
}
