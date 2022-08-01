package service

import (
	"Parrillada/internal/cocina/repository"
	"Parrillada/models/comensal"
	"Parrillada/models/comida"
)

//go:generate mockgen -source=service.go -destination mock_comensales_service_test.go -package service

type ComensalesRepository interface {
	AgregarComensal(comensal comensal.IComensal) []comensal.IComensal
	ObtenerComidasPreferidas(id string) []comida.IComida
	ElegirPlato(id string, comidasDisponibles []comida.IComida) string
	LeAgradaPlato(plato comida.IComida, id string) []comida.IComida
}

type Service struct {
	comensalesRepository repository.Repository
}

func NewComensalesService(comensalesRepository repository.Repository) Service {
	return Service{comensalesRepository}
}

func (s *Service) AgregarComensal(comensal comensal.IComensal) []comensal.IComensal {
	return s.comensalesRepository.AgregarComensal(comensal)
}

func (s *Service) ObtenerComidasPreferidas(id string) []comida.IComida {
	return s.comensalesRepository.ObtenerComidasPreferidas(id)
}

func (s *Service) ElegirPlato(id string, comidasDisponibles []comida.IComida) string {
	return s.comensalesRepository.ElegirPlato(id, comidasDisponibles)
}
func (s *Service) LeAgradaPlato(plato comida.IComida, id string) []comida.IComida {
	return s.comensalesRepository.LeAgradaPlato(plato, id)
}
