package comensal

import "Parrillada/models/comida"

type IComensal interface {
	LeAgradaComida(comida comida.Comida) bool
	Satisfecho() bool
}

type Comensal struct {
	ComidasFavoritas  []comida.Comida `json:"comidas_favoritas"`
	ComidasConsumidas []comida.Comida `json:"comidas_consumidas"`
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
