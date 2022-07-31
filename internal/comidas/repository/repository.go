package repository

import (
	"Parrillada/models/comida"
	"reflect"
)

type Repository struct {
	Comidas []comida.IComida
}

func NewComidaRepository(comidas []comida.IComida) Repository {
	return Repository{Comidas: comidas}
}

func (r *Repository) AgregarPlato(plato comida.IComida) []comida.IComida {
	r.Comidas = append(r.Comidas, plato)
	return r.Comidas
}

func (r *Repository) EliminarPlato(id string) []comida.IComida {
	for index, c := range r.Comidas {
		switch RetornarStructHijoNombre(c) {
		case "*comida.Provoleta":
			p := c.(*comida.Provoleta)
			if p.Comida.ID == id {
				// Remove the element at index index from comidas.
				r.Comidas[index] = r.Comidas[len(r.Comidas)-1] // Copy last element to index i.
				r.Comidas[len(r.Comidas)-1] = nil              // Erase last element (write zero value).
				r.Comidas = r.Comidas[:len(r.Comidas)-1]       // Truncate slice.
			}
		case "*comida.HamburguesaVegana":
			hv := c.(*comida.HamburguesaVegana)
			if hv.Hamburgesa.Comida.ID == id {
				r.Comidas[index] = r.Comidas[len(r.Comidas)-1]
				r.Comidas[len(r.Comidas)-1] = nil
				r.Comidas = r.Comidas[:len(r.Comidas)-1]
			}
		case "*comida.HamburguesaDeCarne":
			hdc := c.(*comida.HamburguesaDeCarne)
			if hdc.Hamburgesa.Comida.ID == id {
				r.Comidas[index] = r.Comidas[len(r.Comidas)-1]
				r.Comidas[len(r.Comidas)-1] = nil
				r.Comidas = r.Comidas[:len(r.Comidas)-1]
			}
		case "*comida.Parrillada":
			p := c.(*comida.Parrillada)
			if p.Comida.ID == id {
				r.Comidas[index] = r.Comidas[len(r.Comidas)-1]
				r.Comidas[len(r.Comidas)-1] = nil
				r.Comidas = r.Comidas[:len(r.Comidas)-1]
			}
		default:
			return nil
		}
	}
	return r.Comidas
}

func (r *Repository) ConsultarPlatosDisponibles() []comida.IComida {
	return r.Comidas
}

func (r *Repository) OfertaVegetariana() bool {
	cantidadPlatosVegetarianos := 0
	cantidadPlatosCarnivoros := 0
	for _, c := range r.Comidas {
		switch RetornarStructHijoNombre(c) {
		case "*comida.Provoleta":
			p := c.(*comida.Provoleta)
			if p.Comida.AptoVegano {
				cantidadPlatosVegetarianos += 1
			} else {
				cantidadPlatosCarnivoros += 1
			}
		case "*comida.HamburguesaVegana":
			hv := c.(*comida.HamburguesaVegana)
			if hv.Hamburgesa.Comida.AptoVegano {
				cantidadPlatosVegetarianos += 1
			} else {
				cantidadPlatosCarnivoros += 1
			}
		case "*comida.HamburguesaDeCarne":
			hdc := c.(*comida.HamburguesaDeCarne)
			if hdc.Hamburgesa.Comida.AptoVegano {
				cantidadPlatosVegetarianos += 1
			} else {
				cantidadPlatosCarnivoros += 1
			}
		case "*comida.Parrillada":
			p := c.(*comida.Parrillada)
			if p.Comida.AptoVegano {
				cantidadPlatosVegetarianos += 1
			} else {
				cantidadPlatosCarnivoros += 1
			}
		}
	}
	if cantidadPlatosCarnivoros > cantidadPlatosVegetarianos {
		return (cantidadPlatosCarnivoros - cantidadPlatosVegetarianos) <= 2
	} else {
		return cantidadPlatosVegetarianos-cantidadPlatosCarnivoros <= 2
	}
}

func (r *Repository) PlatoCarnivoroMasFuerte() comida.IComida {
	indiceMejorValorado := 0
	mayorValoracion := 0
	for indice, c := range r.Comidas {
		switch RetornarStructHijoNombre(c) {
		case "*comida.Provoleta":
			p := c.(*comida.Provoleta)
			if p.Comida.AptoVegano && p.Comida.Valoracion > mayorValoracion {
				mayorValoracion = p.Comida.Valoracion
				indiceMejorValorado = indice
			}
		case "*comida.HamburguesaDeCarne":
			hdc := c.(*comida.HamburguesaDeCarne)
			if hdc.Hamburgesa.Comida.Valoracion > mayorValoracion {
				mayorValoracion = hdc.Hamburgesa.Comida.Valoracion
				indiceMejorValorado = indice
			}
		case "*comida.Parrillada":
			p := c.(*comida.Parrillada)
			if p.Comida.AptoVegano && p.Comida.Valoracion > mayorValoracion {
				mayorValoracion = p.Comida.Valoracion
				indiceMejorValorado = indice
			}
		}
	}
	return r.Comidas[indiceMejorValorado]
}

func RetornarStructHijoNombre(plato comida.IComida) string {

	childStruct := reflect.TypeOf(plato)

	return childStruct.String()
}
