package comida

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

/*
	Tests requeridos
*/
func TestProvoletaEspecial(t *testing.T) {
	p := NewProvoleta(190.0, true)

	assert.Equal(t, true, p.Especial)
}

/*
 Test no rquerido
*/

func TestCalcularValoracionEspecial(t *testing.T) {
	pe := NewProvoleta(190.0, true)
	pe.CalcularValoracion()

	assert.Equal(t, 120, pe.Comida.Valoracion)
}

func TestCalcularValoracionComun(t *testing.T) {
	pe := NewProvoleta(190.0, false)
	pe.CalcularValoracion()

	assert.Equal(t, 80, pe.Comida.Valoracion)
}

func TestCalcularValoracionEspecialPeso(t *testing.T) {
	pe := NewProvoleta(251.0, false)
	pe.CalcularValoracion()

	assert.Equal(t, 120, pe.Comida.Valoracion)
}
