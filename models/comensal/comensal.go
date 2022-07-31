package comensal

import (
	"Parrillada/models/comida"
)

type IComensal interface {
	LeAgradaComida(comida comida.IComida) bool
	Satisfecho() bool
}

type Comensal struct {
	Peso              float64          `json:"peso"`
	ComidasFavoritas  []comida.IComida `json:"comidas_favoritas"`
	ComidasConsumidas []comida.IComida `json:"comidas_consumidas"`
}

func (c *Comensal) comer(comida comida.IComida) {
	c.ComidasConsumidas = append(c.ComidasConsumidas, comida)
}
