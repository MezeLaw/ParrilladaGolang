package handler

import (
	"Parrillada/internal/cocina/service"
	"Parrillada/models/comensal"
	"Parrillada/models/comida"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type ComensalesService interface {
	AgregarComensal(comensal comensal.IComensal) []comensal.IComensal
	ObtenerComidasPreferidas(id string) []comida.IComida
	ElegirPlato(id string) string
	LeAgradaPlato(plato comida.IComida, id string) []comida.IComida
}

type Handler struct {
	comensalesService service.Service
}

func NewComensalHandler(cs service.Service) Handler {
	return Handler{comensalesService: cs}
}

func (h *Handler) AgregarComensal(c *gin.Context) {
	comensl := comensal.Comensal{}

	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Comensal invalido!")
		return
	}

	json.Unmarshal(jsonData, &comensl)

	if comensl.TipoDeComensal == "vegetariano" {
		vegetariano := comensal.NewVegetariano(comensl.Peso, comensl.TipoDeComensal)
		resp := h.comensalesService.AgregarComensal(&vegetariano)
		if len(resp) < 1 {
			c.JSON(http.StatusInternalServerError, "Ocurrio un error al intentar agregar el comensal")
			return
		}
		c.JSON(http.StatusCreated, resp)
		return
	}

	if comensl.TipoDeComensal == "hambrePopular" {
		hp := comensal.NewHambrePopular(comensl.Peso, comensl.TipoDeComensal)
		resp := h.comensalesService.AgregarComensal(&hp)
		if len(resp) < 1 {
			c.JSON(http.StatusInternalServerError, "Ocurrio un error al intentar agregar el comensal")
			return
		}
		c.JSON(http.StatusCreated, resp)
		return
	}

	if comensl.TipoDeComensal == "paladarFino" {
		pf := comensal.NewPaladarFino(comensl.Peso, comensl.TipoDeComensal)
		resp := h.comensalesService.AgregarComensal(&pf)
		if len(resp) < 1 {
			c.JSON(http.StatusInternalServerError, "Ocurrio un error al intentar agregar el comensal")
			return
		}
		c.JSON(http.StatusCreated, resp)
		return
	}
	c.JSON(http.StatusBadRequest, "El comensal no es valido!")
	return
}

func (h *Handler) ObtenerComidasPreferidas(c *gin.Context) {
	comensalId := c.Param("id")

	if comensalId == "" {
		c.JSON(http.StatusBadRequest, "Ingrese un id de comensal valido")
	}

	resp := h.comensalesService.ObtenerComidasPreferidas(comensalId)

	if resp == nil {
		c.JSON(http.StatusNotFound, "No se encontro el comensal")
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}

func (h *Handler) ElegirPlato(c *gin.Context) {
	comensalId := c.Param("id")

	if comensalId == "" {
		c.JSON(http.StatusBadRequest, "Ingrese un id de comensal valido")
	}

	resp := h.comensalesService.ElegirPlato(comensalId)

	c.JSON(http.StatusOK, resp)
	return
}

func (h *Handler) LeAgradaPlato(c *gin.Context) {
	comensalId := c.Param("id")

	if comensalId == "" {
		c.JSON(http.StatusBadRequest, "Ingrese un id de comensal valido")
	}

	//asdasdasd
	provoleta := comida.Provoleta{}
	hamburguesaVegana := comida.HamburguesaVegana{}
	hamburguesaDeCarne := comida.HamburguesaDeCarne{}
	parrillada := comida.Parrillada{}

	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Comida invalida!")
		return
	}

	json.Unmarshal(jsonData, &hamburguesaVegana)
	json.Unmarshal(jsonData, &hamburguesaDeCarne)
	json.Unmarshal(jsonData, &parrillada)
	json.Unmarshal(jsonData, &provoleta)

	if hamburguesaVegana.IngredienteMedallon != "" {

		hv := comida.NewHamburguesaVegana(hamburguesaVegana.Hamburgesa.TipoDePan, hamburguesaVegana.IngredienteMedallon)
		resp := h.comensalesService.LeAgradaPlato(&hv, comensalId)
		if resp != nil {
			c.JSON(http.StatusCreated, resp)
			return
		}
		c.JSON(http.StatusInternalServerError, "Al comensal no le gusta el plato")
		return
	}

	if hamburguesaDeCarne.Hamburgesa.TipoDePan != "" {
		hdc := comida.NewHamburguesaDeCarne(hamburguesaDeCarne.Hamburgesa.TipoDePan)
		resp := h.comensalesService.LeAgradaPlato(&hdc, comensalId)
		if resp != nil {
			c.JSON(http.StatusCreated, resp)
			return
		}
		c.JSON(http.StatusInternalServerError, "No se pudo agregar la hamburguesa de carne")
		return
	}

	if len(parrillada.Cortes) > 0 {
		parri := comida.NewParrillada(parrillada.Cortes)
		resp := h.comensalesService.LeAgradaPlato(&parri, comensalId)
		if resp != nil {
			c.JSON(http.StatusCreated, resp)
			return
		}
		c.JSON(http.StatusInternalServerError, "No se pudo agregar la parrillada")
		return
	}

	if provoleta.Peso > 0 {
		provo := comida.NewProvoleta(provoleta.Comida.Peso, provoleta.TieneEspecias)
		resp := h.comensalesService.LeAgradaPlato(&provo, comensalId)
		if resp != nil {
			c.JSON(http.StatusCreated, resp)
			return
		} else {
			c.JSON(http.StatusInternalServerError, "No se pudo agregar la provoleta")
			return
		}
	}

	c.JSON(http.StatusBadRequest, "No se reconocio la comida enviada")
	return
}
