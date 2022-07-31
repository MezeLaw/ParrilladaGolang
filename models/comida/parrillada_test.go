package comida

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

/*
	Test requerido
*/
func TestValoracionParrillada(t *testing.T) {
	/*
		Parillada, que incluya un asado, de calidad 10 y peso 250 gramos, una entraña,
		de calidad 7 y peso 200 gramos y un chorizo, de calidad 8 y peso 50 gramos.
		El resultado de la valoracion será de 147 (15 * 10) - 3.
	*/

	asado := NewCorte("asado", 10, 250)
	entrania := NewCorte("entrania", 7, 200)
	chori := NewCorte("chorizo", 8, 50)

	var cortes []Corte
	cortes = append(cortes, asado)
	cortes = append(cortes, entrania)
	cortes = append(cortes, chori)

	parrillada := NewParrillada(cortes)
	parrillada.CalcularValoracion()

	assert.Equal(t, 147, parrillada.Comida.Valoracion)
}

/*
	Test adicionales
*/
func TestValoracionParrilladaConCorteCalidadInvalida(t *testing.T) {
	/*
		Parillada, que incluya un asado, de calidad 10 y peso 250 gramos, una entraña,
		de calidad 7 y peso 200 gramos y un chorizo, de calidad 8 y peso 50 gramos.
		El resultado de la valoracion será de 147 (15 * 10) - 3.
	*/

	asado := NewCorte("asado", 10, 250)
	entrania := NewCorte("entrania", 7, 200)
	chori := NewCorte("chorizo", 0, 50)

	var cortes []Corte
	cortes = append(cortes, asado)
	cortes = append(cortes, entrania)
	cortes = append(cortes, chori)

	parrillada := NewParrillada(cortes)
	parrillada.CalcularValoracion()

	assert.Equal(t, 147, parrillada.Comida.Valoracion)
}

func TestValoracionParrilladaInvalida(t *testing.T) {
	/*
		Parillada, que incluya un asado, de calidad 10 y peso 250 gramos, una entraña,
		de calidad 7 y peso 200 gramos y un chorizo, de calidad 8 y peso 50 gramos.
		El resultado de la valoracion será de 147 (15 * 10) - 3.
	*/

	asado := NewCorte("asado", 0, 250)
	entrania := NewCorte("entrania", 0, 200)
	chori := NewCorte("chorizo", 0, 50)

	asado2 := NewCorte("asado", 0, 250)
	entrania2 := NewCorte("entrania", 0, 200)
	chori2 := NewCorte("chorizo", 0, 50)

	asado3 := NewCorte("asado", 0, 250)
	entrania3 := NewCorte("entrania", 0, 200)
	chori3 := NewCorte("chorizo", 0, 50)

	asado4 := NewCorte("asado", 0, 250)
	entrania4 := NewCorte("entrania", 0, 200)
	chori4 := NewCorte("chorizo", 0, 50)

	asado5 := NewCorte("asado", 0, 250)
	entrania5 := NewCorte("entrania", 0, 200)
	chori5 := NewCorte("chorizo", 0, 50)
	var cortes []Corte

	cortes = append(cortes, asado)
	cortes = append(cortes, entrania)
	cortes = append(cortes, chori)

	cortes = append(cortes, asado2)
	cortes = append(cortes, entrania2)
	cortes = append(cortes, chori2)

	cortes = append(cortes, asado3)
	cortes = append(cortes, entrania3)
	cortes = append(cortes, chori3)

	cortes = append(cortes, asado4)
	cortes = append(cortes, entrania4)
	cortes = append(cortes, chori4)

	cortes = append(cortes, asado5)
	cortes = append(cortes, entrania5)
	cortes = append(cortes, chori5)

	parrillada := NewParrillada(cortes)
	parrillada.CalcularValoracion()

	assert.Equal(t, 1, parrillada.Comida.Valoracion)
}
