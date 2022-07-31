package comida

type IComida interface {
	CalcularValoracion()
}

type Comida struct {
	ID         string  `json:"id"`
	Peso       float64 `json:"peso"`
	AptoVegano bool    `json:"apto_vegano"`
	Valoracion int     `json:"valoracion"`
}

func (c *Comida) esAbundante() bool {
	return c.Peso > 250
}
