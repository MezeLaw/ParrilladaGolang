package repository

import (
	"Parrillada/models/comida"
	"testing"
)

func TestRetornarStructHijoNombres(t *testing.T) {
	comidas := []comida.IComida{}

	p := comida.NewProvoleta(190, true)
	hv := comida.NewHamburguesaVegana(comida.MasaMadre, "lentejas")
	hdc := comida.NewHamburguesaDeCarne(comida.MasaMadre)

	cortes := []comida.Corte{comida.NewCorte("asado", 10, 150)}
	parrillada := comida.NewParrillada(cortes)

	comidas = append(comidas, &p)
	comidas = append(comidas, &hv)
	comidas = append(comidas, &hdc)
	comidas = append(comidas, &parrillada)

	for _, c := range comidas {
		println(RetornarStructHijoNombre(c))
	}

	/*
	*comida.Provoleta
	*comida.HamburguesaVegana
	*comida.HamburguesaDeCarne
	*comida.Parrillada
	 */
}
