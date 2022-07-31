// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service is a generated GoMock package.
package service

import (
	comida "Parrillada/models/comida"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockComidasRepository is a mock of ComidasRepository interface.
type MockComidasRepository struct {
	ctrl     *gomock.Controller
	recorder *MockComidasRepositoryMockRecorder
}

// MockComidasRepositoryMockRecorder is the mock recorder for MockComidasRepository.
type MockComidasRepositoryMockRecorder struct {
	mock *MockComidasRepository
}

// NewMockComidasRepository creates a new mock instance.
func NewMockComidasRepository(ctrl *gomock.Controller) *MockComidasRepository {
	mock := &MockComidasRepository{ctrl: ctrl}
	mock.recorder = &MockComidasRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComidasRepository) EXPECT() *MockComidasRepositoryMockRecorder {
	return m.recorder
}

// AgregarPlato mocks base method.
func (m *MockComidasRepository) AgregarPlato(plato comida.IComida) []comida.IComida {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgregarPlato", plato)
	ret0, _ := ret[0].([]comida.IComida)
	return ret0
}

// AgregarPlato indicates an expected call of AgregarPlato.
func (mr *MockComidasRepositoryMockRecorder) AgregarPlato(plato interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgregarPlato", reflect.TypeOf((*MockComidasRepository)(nil).AgregarPlato), plato)
}

// ConsultarPlatosDisponibles mocks base method.
func (m *MockComidasRepository) ConsultarPlatosDisponibles() []comida.IComida {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsultarPlatosDisponibles")
	ret0, _ := ret[0].([]comida.IComida)
	return ret0
}

// ConsultarPlatosDisponibles indicates an expected call of ConsultarPlatosDisponibles.
func (mr *MockComidasRepositoryMockRecorder) ConsultarPlatosDisponibles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsultarPlatosDisponibles", reflect.TypeOf((*MockComidasRepository)(nil).ConsultarPlatosDisponibles))
}

// EliminarPlato mocks base method.
func (m *MockComidasRepository) EliminarPlato(id string) []comida.IComida {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EliminarPlato", id)
	ret0, _ := ret[0].([]comida.IComida)
	return ret0
}

// EliminarPlato indicates an expected call of EliminarPlato.
func (mr *MockComidasRepositoryMockRecorder) EliminarPlato(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EliminarPlato", reflect.TypeOf((*MockComidasRepository)(nil).EliminarPlato), id)
}

// OfertaVegetariana mocks base method.
func (m *MockComidasRepository) OfertaVegetariana() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfertaVegetariana")
	ret0, _ := ret[0].(bool)
	return ret0
}

// OfertaVegetariana indicates an expected call of OfertaVegetariana.
func (mr *MockComidasRepositoryMockRecorder) OfertaVegetariana() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfertaVegetariana", reflect.TypeOf((*MockComidasRepository)(nil).OfertaVegetariana))
}

// PlatoCarnivoroMasFuerte mocks base method.
func (m *MockComidasRepository) PlatoCarnivoroMasFuerte() comida.IComida {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlatoCarnivoroMasFuerte")
	ret0, _ := ret[0].(comida.IComida)
	return ret0
}

// PlatoCarnivoroMasFuerte indicates an expected call of PlatoCarnivoroMasFuerte.
func (mr *MockComidasRepositoryMockRecorder) PlatoCarnivoroMasFuerte() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlatoCarnivoroMasFuerte", reflect.TypeOf((*MockComidasRepository)(nil).PlatoCarnivoroMasFuerte))
}