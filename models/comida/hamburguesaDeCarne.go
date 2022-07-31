package comida

type HamburguesaDeCarne struct {
	Hamburgesa
}

func NewHamburguesaDeCarne(tipoDePan string) HamburguesaDeCarne {

	h := NewHamburguesa(tipoDePan)
	h.Comida.AptoVegano = false

	hdc := HamburguesaDeCarne{
		Hamburgesa: h,
	}
	hdc.CalcularValoracion()
	return hdc
}

func (hdc *HamburguesaDeCarne) CalcularValoracion() {
	hdc.Hamburgesa.CalcularValoracionPorTipoDePan()
}
