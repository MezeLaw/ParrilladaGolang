package comensal

import (
	"Parrillada/models/comida"
	"Parrillada/utils"
	"github.com/google/uuid"
	"log"
)

type Vegetariano struct {
	Comensal
}

func NewVegetariano(peso float64, tipoDeComensal string) Vegetariano {
	comensal := Comensal{
		ID:                uuid.New().String(),
		Peso:              peso,
		ComidasFavoritas:  []comida.IComida{},
		ComidasConsumidas: []comida.IComida{},
		TipoDeComensal:    tipoDeComensal,
	}

	return Vegetariano{comensal}
}

func (v *Vegetariano) LeAgradaComida(c comida.IComida) bool {

	//Todas con calific > 85 & aptoVeg
	switch utils.RetornarNombreStructHijoComida(c) {
	case "*comida.Provoleta":
		p := c.(*comida.Provoleta)
		if p.Comida.AptoVegano && p.Comida.Valoracion > 85 {
			v.ComidasFavoritas = append(v.ComidasFavoritas, p)
			return true
		}

	case "*comida.HamburguesaVegana":
		hv := c.(*comida.HamburguesaVegana)
		if hv.Comida.Valoracion > 85 {
			v.ComidasFavoritas = append(v.ComidasFavoritas, hv)
			return true
		}

	default:
		return false
	}
	return false
}

func (v *Vegetariano) Satisfecho() bool {
	pesoComidaConsumida := 0.0

	for _, c := range v.ComidasConsumidas {

		switch utils.RetornarNombreStructHijoComida(c) {
		case "*comida.Provoleta":
			p := c.(*comida.Provoleta)
			pesoComidaConsumida += p.Peso

		case "*comida.HamburguesaVegana":
			hv := c.(*comida.HamburguesaVegana)
			pesoComidaConsumida += hv.Peso
		case "*comida.HamburguesaDeCarne":
			hdc := c.(*comida.HamburguesaDeCarne)
			pesoComidaConsumida += hdc.Peso
		case "*comida.Parrillada":
			p := c.(*comida.Parrillada)
			pesoComidaConsumida += p.Peso
		default:
			log.Println("No se pudo reconocer la comida. No sera tenida en cuenta para medir satisfaccion")
			pesoComidaConsumida += 0
		}

	}
	return pesoComidaConsumida >= v.Peso*0.01 && v.NingunaEsAbundante()
}

func (v *Vegetariano) NingunaEsAbundante() bool {

	for _, c := range v.ComidasConsumidas {

		switch utils.RetornarNombreStructHijoComida(c) {
		case "*comida.Provoleta":
			p := c.(*comida.Provoleta)
			if p.Comida.EsAbundante() {
				return false
			}

		case "*comida.HamburguesaVegana":
			hv := c.(*comida.HamburguesaVegana)
			if hv.Comida.EsAbundante() {
				return false
			}
		case "*comida.HamburguesaDeCarne":
			hdc := c.(*comida.HamburguesaDeCarne)
			if hdc.Hamburgesa.Comida.EsAbundante() {
				return false
			}
		case "*comida.Parrillada":
			p := c.(*comida.Parrillada)
			if p.Comida.EsAbundante() {
				return false
			}
		}
	}

	return true
}
