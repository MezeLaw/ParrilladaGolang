package service

import (
	"Parrillada/models/comida"
)

//go:generate mockgen -source=service.go -destination mock_comidas_service_test.go -package service

type ComidasRepository interface {
	AgregarPlato(plato comida.IComida) []comida.IComida
	EliminarPlato(id string) []comida.IComida
	ConsultarPlatosDisponibles() []comida.IComida
	OfertaVegetariana() bool
	PlatoCarnivoroMasFuerte() comida.IComida
}

type Service struct {
	comidasRepository ComidasRepository
}

func NewComidaService(comidasRepository ComidasRepository) Service {
	return Service{comidasRepository: comidasRepository}
}

func (s *Service) AgregarPlato(plato comida.IComida) []comida.IComida {
	return s.comidasRepository.AgregarPlato(plato)
}

func (s *Service) EliminarPlato(id string) []comida.IComida {
	return s.comidasRepository.EliminarPlato(id)
}

func (s *Service) ConsultarPlatosDisponibles() []comida.IComida {
	return s.comidasRepository.ConsultarPlatosDisponibles()
}

func (s *Service) OfertaVegetariana() bool {
	return s.comidasRepository.OfertaVegetariana()
}

func (s *Service) PlatoCarnivoroMasFuerte() comida.IComida {
	return s.comidasRepository.PlatoCarnivoroMasFuerte()
}
