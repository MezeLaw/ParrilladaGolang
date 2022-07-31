package comida

type HamburguesaVegana struct {
	Hamburgesa
	IngredienteMedallon string `json:"ingrediente_medallon"`
}

func NewHamburguesaVegana(tipoDePan string, ingredienteMedallon string) HamburguesaVegana {
	h := NewHamburguesa(tipoDePan)
	h.Comida.AptoVegano = true
	hv := HamburguesaVegana{
		Hamburgesa:          h,
		IngredienteMedallon: ingredienteMedallon,
	}
	hv.CalcularValoracion()
	return hv
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
