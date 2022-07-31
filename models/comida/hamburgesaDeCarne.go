package comida

type HamburguesaDeCarne struct {
	Hamburgesa Hamburgesa `json:"hamburgesa"`
}

func NewHamburguesaDeCarne(tipoDePan string) IComida {

	h := NewHamburguesa(tipoDePan)
	h.Comida.AptoVegano = false

	hdc := HamburguesaDeCarne{
		Hamburgesa: h,
	}
	return &hdc
}

func (hdc *HamburguesaDeCarne) CalcularValoracion() {
	hdc.Hamburgesa.CalcularValoracionPorTipoDePan()
}
