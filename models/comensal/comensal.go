package comensal

import "Parrillada/models/comida"

type IComensal interface {
	LeAgradaComida(comida comida.Comida) bool
	Satisfecho() bool
}

type Comensal struct {
	ComidasFavoritas  []comida.Comida
	ComidasConsumidas []comida.Comida
}

func NewComensal(comidasFavoritas []comida.Comida, comidasConsumidas []comida.Comida) Comensal {
	return Comensal{
		ComidasFavoritas:  comidasFavoritas,
		ComidasConsumidas: comidasConsumidas,
	}
}

func (c *Comensal) comer(comida comida.Comida) {
	c.ComidasConsumidas = append(c.ComidasConsumidas, comida)
}
