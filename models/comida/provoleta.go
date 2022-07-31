package comida

import "github.com/google/uuid"

type Provoleta struct {
	Comida        Comida `json:"comida"`
	TieneEspecias bool   `json:"tiene_especias"`
	Especial      bool   `json:"especial"`
}

func NewProvoleta(peso float64, tieneEspecias bool) Provoleta {
	c := Comida{
		Peso:       peso,
		AptoVegano: !tieneEspecias,
		ID:         uuid.New().String(),
	}
	return Provoleta{
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
