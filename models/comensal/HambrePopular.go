package comensal

import (
	"Parrillada/models/comida"
	"Parrillada/utils"
	"github.com/google/uuid"
	"log"
)

type HambrePopular struct {
	Comensal
}

func NewHambrePopular(peso float64, tipoDeComensal string) HambrePopular {

	comidasFavoritas := []comida.IComida{}
	comidasConsumidas := []comida.IComida{}

	comensal := Comensal{
		ID:                uuid.New().String(),
		Peso:              peso,
		ComidasFavoritas:  comidasFavoritas,
		ComidasConsumidas: comidasConsumidas,
		TipoDeComensal:    tipoDeComensal,
	}

	return HambrePopular{Comensal: comensal}
}

func (hp *HambrePopular) LeAgradaComida(c comida.IComida) bool {

	switch utils.RetornarNombreStructHijoComida(c) {
	case "*comida.Provoleta":
		p := c.(*comida.Provoleta)
		if p.Comida.EsAbundante() {
			hp.ComidasFavoritas = append(hp.ComidasFavoritas, c)
			return true
		}

	case "*comida.HamburguesaVegana":
		hv := c.(*comida.HamburguesaVegana)
		if hv.Comida.EsAbundante() {
			//Siempre deberia NO SER abundante, pero por las dudas pregunto
			hp.ComidasFavoritas = append(hp.ComidasFavoritas, c)
			return true
		}
	case "*comida.HamburguesaDeCarne":
		hdc := c.(*comida.HamburguesaDeCarne)
		if hdc.Hamburgesa.Comida.EsAbundante() {
			hp.ComidasFavoritas = append(hp.ComidasFavoritas, c)
			return true
		}
	case "*comida.Parrillada":
		p := c.(*comida.Parrillada)
		if p.Comida.EsAbundante() {
			hp.ComidasFavoritas = append(hp.ComidasFavoritas, c)
			return true
		}
	default:
		return false
	}
	return false
}

func (hp *HambrePopular) Satisfecho() bool {

	pesoComidaConsumida := 0.0

	for _, c := range hp.ComidasConsumidas {

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
	return pesoComidaConsumida >= hp.Peso*0.01
}
