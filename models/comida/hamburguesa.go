package comida

import "github.com/google/uuid"

type Hamburgesa struct {
	Comida         Comida `json:"comida"`
	ValoracionBase int    `json:"valoracion_base"`
	TipoDePan      string `json:"tipo_de_pan"`
}

const (
	PanCasero  = "casero"
	MasaMadre  = "masaMadre"
	Industrial = "industrial"
)

func NewHamburguesa(tipoDePan string) Hamburgesa {

	comida := Comida{
		Peso: 250,
		ID:   uuid.New().String(),
	}
	//Por default si el pan es invalido construyo la hamburgesa con industrial
	if PanValido(tipoDePan) {
		return Hamburgesa{
			Comida:         comida,
			ValoracionBase: 60,
			TipoDePan:      tipoDePan,
		}
	}
	return Hamburgesa{
		Comida:         comida,
		ValoracionBase: 60,
		TipoDePan:      Industrial,
	}
}

func PanValido(tipoDePan string) bool {
	if tipoDePan == PanCasero || tipoDePan == MasaMadre || tipoDePan == Industrial {
		return true
	}
	return false
}

func (h *Hamburgesa) CalcularValoracionPorTipoDePan() {
	if h.TipoDePan == MasaMadre {
		h.Comida.Valoracion = h.ValoracionBase + 45
	}

	if h.TipoDePan == PanCasero {
		h.Comida.Valoracion = h.ValoracionBase + 20
	}

	if h.TipoDePan == Industrial {
		h.Comida.Valoracion = h.ValoracionBase
	}
}
