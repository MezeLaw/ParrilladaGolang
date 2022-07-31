package comida

type Provoleta struct {
	Comida        Comida `json:"comida"`
	TieneEspecias bool   `json:"tiene_especias"`
	Especial      bool   `json:"especial"`
}

func NewProvoleta(peso float64, tieneEspecias bool) IComida {
	c := Comida{
		Peso:       peso,
		AptoVegano: !tieneEspecias,
	}
	return &Provoleta{
		Comida:        c,
		TieneEspecias: tieneEspecias,
		Especial:      tieneEspecias || c.esAbundante(),
	}
}

func (p *Provoleta) CalcularValoracion() {
	if p.Especial {
		p.Comida.Valoracion = 120
	} else {
		p.Comida.Valoracion = 80
	}
}
