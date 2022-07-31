package main

import (
	"Parrillada/internal/comidas/handler"
	"Parrillada/internal/comidas/repository"
	"Parrillada/internal/comidas/service"
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

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	router.POST("/cocina/platos/", comidasHandler.AgregarPlato)
	router.GET("/cocina/platos", comidasHandler.ConsultarPlatosDisponibles)
	router.GET("/cocina/ofertaVegetariana", comidasHandler.OfertaVegetariana)
	router.GET("/cocina/platoCarnivoroMasFuerte", comidasHandler.PlatoCarnivoroMasFuerte)

	router.Run()
}
