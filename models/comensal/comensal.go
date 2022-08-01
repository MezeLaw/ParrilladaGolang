package comensal

import (
	"Parrillada/models/comida"
)

type IComensal interface {
	LeAgradaComida(comida comida.IComida) bool
	Satisfecho() bool
}

type Comensal struct {
	ID                string           `json:"id"`
	Peso              float64          `json:"peso"`
	ComidasFavoritas  []comida.IComida `json:"comidas_favoritas"`
	ComidasConsumidas []comida.IComida `json:"comidas_consumidas"`
	TipoDeComensal    string           `json:"tipo_de_comensal"`
}

func (c *Comensal) Comer(comida comida.IComida) {
	c.ComidasConsumidas = append(c.ComidasConsumidas, comida)
}
