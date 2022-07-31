package comensal

import (
	"Parrillada/models/comida"
	"github.com/go-playground/assert/v2"
	"testing"
)

/**
Tests requeridos
*/

/*
Un vegetariano de peso 68500 gramos come
una provoleta de 190 gramos con especias y
dos Hamburguesas Vegetarianas (250 grs cada una) con pan de masaMadre y de garbanzos como legumbre.
El resultado si esta satisfecho es: True porque (190 + 250 + 250) >= (68500 * 0.01)
y además ninguna de las tres comidas es abundante, es decir, ninguna supera los 250 gramos.
*/

func TestVegetarianoSatisfecho(t *testing.T) {
	vegetariano := NewVegetariano(68500)

	provoleta := comida.NewProvoleta(190, true)
	hv1 := comida.NewHamburguesaVegana(comida.MasaMadre, "garbanzos")
	hv2 := comida.NewHamburguesaVegana(comida.MasaMadre, "garbanzos")

	vegetariano.Comer(&provoleta)
	vegetariano.Comer(&hv1)
	vegetariano.Comer(&hv2)

	satisfecho := vegetariano.Satisfecho()

	assert.Equal(t, true, satisfecho)
}

/*
Un comensal popular de peso 85000 gramos come
una parrillada que incluye
un asado, de calidad 10 y peso 250 gramos,
una entraña, de calidad 7 y peso 200 gramos
y un chorizo, de calidad 8 y peso 50 gramos.
El resultado si esta satisfecho es: False porque 500 < 850
*/

func TestComensalPopularInsatisfecho(t *testing.T) {
	cp := NewHambrePopular(85000)

	asado := comida.NewCorte("asado", 10, 250)
	entrania := comida.NewCorte("entrania", 7, 200)
	chori := comida.NewCorte("chorizo", 8, 50)

	cortes := []comida.Corte{asado, entrania, chori}

	parrillada := comida.NewParrillada(cortes)

	cp.Comer(&parrillada)

	satisfecho := cp.Satisfecho()

	assert.Equal(t, false, satisfecho)

}

/*
	Un comensal de paladar fino le agrada comer una Hamburguesa de carne de 250 gramos con pan de masa madre.
*/

func TestComensalPaladarFinoLeAgradaHamburguesaDeMasaMadre(t *testing.T) {
	pf := NewPaladarFino(85000)
	hc := comida.NewHamburguesaDeCarne(comida.MasaMadre)

	leAgrada := pf.LeAgradaComida(&hc)

	assert.Equal(t, true, leAgrada)
}

/*
	Un comensal de paladar fino NO agrada comer una Hamburguesa de carne de 250 gramos con pan casero.
*/

func TestComensalPaladarFinoNoLeAgradaHamburguesaDeCarneConPanCasero(t *testing.T) {
	pf := NewPaladarFino(85000)
	hc := comida.NewHamburguesaDeCarne(comida.PanCasero)

	leAgrada := pf.LeAgradaComida(&hc)

	assert.Equal(t, false, leAgrada)
}
