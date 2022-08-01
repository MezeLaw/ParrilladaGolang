package comida

import "github.com/google/uuid"

type Parrillada struct {
	Comida
	Cortes []Corte `json:"cortes"`
}

func NewParrillada(cortes []Corte) Parrillada {
	pesoParrillada := CalcularPesoDeParrillada(cortes)
	c := Comida{
		Peso:       pesoParrillada,
		AptoVegano: false,
		Valoracion: CalcularValoracionParrillada(cortes),
		ID:         uuid.New().String(),
	}
	parrillada := Parrillada{
		Comida: c,
		Cortes: cortes,
	}
	parrillada.CalcularValoracion()
	parrillada.Comida.Peso = CalcularPesoDeParrillada(parrillada.Cortes)
	return parrillada
}

func CalcularPesoDeParrillada(cortes []Corte) float64 {
	pesoParrillada := 0.0
	for _, corte := range cortes {
		pesoParrillada += corte.Peso
	}
	return pesoParrillada
}

func CalcularValoracionParrillada(cortes []Corte) int {
	valoracion := (15 * CalcularCorteConMayorValoracion(cortes)) - len(cortes)
	if valoracion < 1 {
		//Lo defino yo, para evitar valoraciones negativas
		return 1
	}
	return valoracion
}

func (p *Parrillada) CalcularValoracion() {
	p.Comida.Valoracion = CalcularValoracionParrillada(p.Cortes)
}

func CalcularCorteConMayorValoracion(cortes []Corte) int {
	mayorValoracion := 1

	for _, corte := range cortes {
		if corte.Calidad > mayorValoracion {
			mayorValoracion = corte.Calidad
		}
	}
	return mayorValoracion
}
