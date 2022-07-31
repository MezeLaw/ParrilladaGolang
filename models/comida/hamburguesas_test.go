package comida

import (
	"github.com/go-playground/assert/v2"
	"reflect"
	"testing"
)

func TestHamburguesaDeCarne_CalcularValoracion(t *testing.T) {
	type fields struct {
		Hamburgesa Hamburgesa
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Test valoracion pan de MasaMadre",
			fields: fields{Hamburgesa: NewHamburguesa(MasaMadre)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hdc := &HamburguesaDeCarne{
				Hamburgesa: tt.fields.Hamburgesa,
			}
			hdc.CalcularValoracion()
		})
	}
}

func TestNewHamburguesaDeCarne(t *testing.T) {
	type args struct {
		tipoDePan string
	}
	tests := []struct {
		name string
		args args
		want HamburguesaDeCarne
	}{
		{
			name: "Test with valid  values should pass",
			args: args{tipoDePan: MasaMadre},
			want: NewHamburguesaDeCarne(MasaMadre),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHamburguesaDeCarne(tt.args.tipoDePan); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHamburguesaDeCarne() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
 * Test requeridos
 */
func TestValoracionHamburguesaPanCasero(t *testing.T) {
	hc := NewHamburguesaDeCarne(PanCasero)

	hc.CalcularValoracion()

	valoracion := hc.Hamburgesa.Comida.Valoracion

	assert.Equal(t, 80, valoracion)
}

func TestValoracionHamburguesaVeganaMasaMadreYGarbanzos(t *testing.T) {
	hv := NewHamburguesaVegana(MasaMadre, "garbanzos")
	hv.CalcularValoracion()

	valoracion := hv.Hamburgesa.Comida.Valoracion

	assert.Equal(t, 122, valoracion)
}

func TestCalcularValoracionHamburguesaDePanVencido(t *testing.T) {

	hc := NewHamburguesaDeCarne("vencido")

	hc.CalcularValoracion()

	assert.Equal(t, 60, hc.Hamburgesa.Comida.Valoracion)
}

func TestValoracionHamburguesaDeMorron(t *testing.T) {

	hv := NewHamburguesaVegana(MasaMadre, "morron")

	hv.CalcularValoracion()

	assert.Equal(t, 117, hv.Hamburgesa.Comida.Valoracion)

}
