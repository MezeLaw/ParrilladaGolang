package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	router.Run()
}
