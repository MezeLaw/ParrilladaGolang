package comida

type Corte struct {
	Nombre  string  `json:"nombre"`
	Calidad int     `json:"calidad"`
	Peso    float64 `json:"peso"`
}

func NewCorte(nombre string, calidad int, peso float64) Corte {

	if !CalidadValida(calidad) {
		peso = 1
	}

	return Corte{
		Nombre:  nombre,
		Calidad: calidad,
		Peso:    peso,
	}
}

//Arbitrariamente defino como calidad minima 1, para evitar algunos escenarios
func CalidadValida(calidad int) bool {
	return calidad > 1
}
