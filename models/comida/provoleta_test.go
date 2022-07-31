package comida

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"reflect"
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

//TODO eliminar este test que es prueba de lo que haremos en la capa web
func TestPrueboUnaListaDeComidasConDifPlatosPuedeSerCreadaYCalcularSusValoraciones(t *testing.T) {
	hv := NewHamburguesaVegana(Industrial, "morron")
	p := NewProvoleta(190, true)

	var platos []IComida
	platos = append(platos, &hv)
	platos = append(platos, &p)

	assert.Equal(t, 2, len(platos))

	platos[1].CalcularValoracion()

	p1 := platos[1]
	fmt.Println(reflect.TypeOf(p1))

	v := p1.(*Provoleta)

	assert.Equal(t, v.TieneEspecias, true)
}
