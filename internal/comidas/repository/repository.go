package repository

import (
	"Parrillada/models/comida"
	"reflect"
)

type Repository struct {
	Comidas []comida.IComida
}

func NewRepository(comidas []comida.IComida) Repository {
	return Repository{Comidas: comidas}
}

func (r *Repository) AgregarPlato(plato comida.IComida) []comida.IComida {
	r.Comidas = append(r.Comidas, plato)
	return r.Comidas
}

/*
func (r *Repository) EliminarPlato(plato comida.IComida) []comida.IComida {
	for _, c := range r.Comidas {

	}
}
*/
func RetornarStructHijoNombre(plato comida.IComida) string {

	childStruct := reflect.TypeOf(plato)

	return childStruct.String()
}
