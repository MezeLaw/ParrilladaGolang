package comensal

import (
	"Parrillada/models/comida"
	"Parrillada/utils"
	"github.com/google/uuid"
	"log"
)

type PaladarFino struct {
	Comensal
}

func NewPaladarFino(peso float64) PaladarFino {
	return PaladarFino{Comensal{
		ID:                uuid.New().String(),
		Peso:              peso,
		ComidasFavoritas:  []comida.IComida{},
		ComidasConsumidas: []comida.IComida{},
	}}
}

func (pf *PaladarFino) LeAgradaComida(c comida.IComida) bool {

	switch utils.RetornarStructHijoNombre(c) {
	case "*comida.Provoleta":
		p := c.(*comida.Provoleta)
		if p.Comida.Peso > 150.0 && p.Comida.Peso < 300.0 && p.Valoracion > 100 {
			pf.ComidasFavoritas = append(pf.ComidasFavoritas, p)
			return true
		}
		return false
	case "*comida.HamburguesaVegana":
		hv := c.(*comida.HamburguesaVegana)
		if hv.Comida.Peso > 150.0 && hv.Comida.Peso < 300.0 && hv.Valoracion > 100 {
			pf.ComidasFavoritas = append(pf.ComidasFavoritas, hv)
			return true
		}
		return false
	case "*comida.HamburguesaDeCarne":
		hdc := c.(*comida.HamburguesaDeCarne)
		if hdc.Comida.Peso > 150.0 && hdc.Comida.Peso < 300.0 && hdc.Valoracion > 100 {
			pf.ComidasFavoritas = append(pf.ComidasFavoritas, hdc)
			return true
		}
		return false
	case "*comida.Parrillada":
		p := c.(*comida.Parrillada)
		if p.Comida.Peso > 150.0 && p.Comida.Peso < 300.0 && p.Valoracion > 100 {
			pf.ComidasFavoritas = append(pf.ComidasFavoritas, p)
			return true
		}
		return false
	default:
		return false
	}
}

func (pf *PaladarFino) Satisfecho() bool {
	pesoComidaConsumida := 0.0

	for _, c := range pf.ComidasConsumidas {

		switch utils.RetornarStructHijoNombre(c) {
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
	return pesoComidaConsumida >= pf.Peso*0.01 && (len(pf.ComidasConsumidas)%2 == 0)
}
