package main

import (
	handler2 "Parrillada/internal/cocina/handler"
	repository2 "Parrillada/internal/cocina/repository"
	service2 "Parrillada/internal/cocina/service"
	"Parrillada/internal/comidas/handler"
	"Parrillada/internal/comidas/repository"
	"Parrillada/internal/comidas/service"
	"Parrillada/models/comensal"
	"Parrillada/models/comida"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	comidas := []comida.IComida{}
	comidasRepository := repository.NewComidaRepository(comidas)
	comidasService := service.NewComidaService(&comidasRepository)
	comidasHandler := handler.NewComidasHandler(comidasService)
	//Comensales
	comensales := []comensal.IComensal{}
	comensalRepository := repository2.NewComensalesRepository(comensales, comidasService)
	comensalService := service2.NewComensalesService(comensalRepository)
	comensalHandler := handler2.NewComensalHandler(comensalService, comidasService)

	//Endpoints
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	router.POST("/cocina/platos/", comidasHandler.AgregarPlato)
	router.GET("/cocina/platos", comidasHandler.ConsultarPlatosDisponibles)
	router.GET("/cocina/ofertaVegetariana", comidasHandler.OfertaVegetariana)
	router.GET("/cocina/platoCarnivoroMasFuerte", comidasHandler.PlatoCarnivoroMasFuerte)
	router.DELETE("/cocina/platos/:id", comidasHandler.EliminarPlato)

	router.POST("/cocina/comensales/", comensalHandler.AgregarComensal)
	router.GET("/cocina/comensales/comidasPreferidas/:id", comensalHandler.ObtenerComidasPreferidas)
	router.GET("/cocina/comensales/elegirPlato/:id", comensalHandler.ElegirPlato)
	router.POST("/cocina/comensales/leAgradaPlato/:id", comensalHandler.LeAgradaPlato)

	router.Run()
}
