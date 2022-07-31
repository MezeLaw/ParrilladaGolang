package comida

type IComida interface {
	CalcularValoracion()
}

type Comida struct {
	Peso       float64 `json:"peso"`
	AptoVegano bool    `json:"apto_vegano"`
	Valoracion int     `json:"valoracion"`
}

func NewComida(peso float64, aptoVegano bool, valoracion int) Comida {
	return Comida{
		Peso:       peso,
		AptoVegano: aptoVegano,
		Valoracion: valoracion,
	}
}

func (c *Comida) esAbundante() bool {
	return c.Peso > 250
}
