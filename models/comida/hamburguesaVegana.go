package comida

type HamburguesaVegana struct {
	Hamburgesa          Hamburgesa `json:"hamburgesa"`
	IngredienteMedallon string     `json:"ingrediente_medallon"`
}

func NewHamburguesaVegana(tipoDePan string, ingredienteMedallon string) IComida {
	h := NewHamburguesa(tipoDePan)
	h.Comida.AptoVegano = true
	hv := HamburguesaVegana{
		Hamburgesa:          h,
		IngredienteMedallon: ingredienteMedallon,
	}
	return &hv
}

func (hv *HamburguesaVegana) CalcularValoracion() {
	hv.Hamburgesa.CalcularValoracionPorTipoDePan()
	hv.Hamburgesa.Comida.Valoracion += hv.CalcularValoracionPorTipoDeMedallon()
}

func (hv *HamburguesaVegana) CalcularValoracionPorTipoDeMedallon() int {
	if len(hv.IngredienteMedallon)*2 > 17 {
		return 17
	}
	return len(hv.IngredienteMedallon) * 2
}
