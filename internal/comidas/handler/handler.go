package handler

import (
	"Parrillada/internal/comidas/service"
	"Parrillada/models/comida"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//go:generate mockgen -source=handler.go -destination mock_comidas_handler_test.go -package handler
type ComidasService interface {
	AgregarPlato(plato comida.IComida) []comida.IComida
	EliminarPlato(id string) []comida.IComida
	ConsultarPlatosDisponibles() []comida.IComida
	OfertaVegetariana() bool
}

type Handler struct {
	comidasService service.Service
}

func NewComidasHandler(service service.Service) Handler {
	return Handler{comidasService: service}
}

func (h *Handler) AgregarPlato(c *gin.Context) {

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
		resp := h.comidasService.AgregarPlato(&hv)
		if resp != nil {
			c.JSON(http.StatusCreated, resp)
			return
		}
		c.JSON(http.StatusInternalServerError, "No se pudo agregar la hamburguesa vegana")
		return
	}

	if hamburguesaDeCarne.Hamburgesa.TipoDePan != "" {
		hdc := comida.NewHamburguesaDeCarne(hamburguesaDeCarne.Hamburgesa.TipoDePan)
		resp := h.comidasService.AgregarPlato(&hdc)
		if resp != nil {
			c.JSON(http.StatusCreated, resp)
			return
		}
		c.JSON(http.StatusInternalServerError, "No se pudo agregar la hamburguesa de carne")
		return
	}

	if len(parrillada.Cortes) > 0 {
		parri := comida.NewParrillada(parrillada.Cortes)
		resp := h.comidasService.AgregarPlato(&parri)
		if resp != nil {
			c.JSON(http.StatusCreated, resp)
			return
		}
		c.JSON(http.StatusInternalServerError, "No se pudo agregar la parrillada")
		return
	}

	if provoleta.Peso > 0 {
		provo := comida.NewProvoleta(provoleta.Comida.Peso, provoleta.TieneEspecias)
		resp := h.comidasService.AgregarPlato(&provo)
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

func (h *Handler) EliminarPlato(c *gin.Context) {
	platoId := c.Param("id")

	if platoId == "" {
		c.JSON(http.StatusBadRequest, "Ingrese un id de plato valido")
	}

	resp := h.comidasService.EliminarPlato(platoId)
	if resp != nil {
		c.JSON(http.StatusNoContent, resp)
		return
	}

	c.JSON(http.StatusInternalServerError, nil)
	return
}

func (h *Handler) ConsultarPlatosDisponibles(c *gin.Context) {
	platos := h.comidasService.ConsultarPlatosDisponibles()

	if platos == nil || len(platos) < 1 {
		c.JSON(http.StatusNotFound, "No hay platos disponibles")
		return
	}

	c.JSON(http.StatusOK, platos)
	return
}

func (h *Handler) OfertaVegetariana(c *gin.Context) {
	platos := h.comidasService.ConsultarPlatosDisponibles()

	if platos == nil || len(platos) < 1 {
		c.JSON(http.StatusNotFound, "No hay platos disponibles")
		return
	}

	resp := h.comidasService.OfertaVegetariana()
	if resp {
		c.JSON(http.StatusOK, "Si tiene buena oferta vegetariana")
		return
	}
	c.JSON(http.StatusOK, "No, no tiene buena oferta vegetariana")
	return
}

func (h *Handler) PlatoCarnivoroMasFuerte(c *gin.Context) {

	res := h.comidasService.PlatoCarnivoroMasFuerte()

	if res == nil {
		c.JSON(http.StatusInternalServerError, "Ocurrio un error al intentar recuperar el plato carnivoro mas valorado")
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
