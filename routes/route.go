package routes

import (
	"net/http"
	"projeto-transportadora/handlers"
	"projeto-transportadora/middlewares"

	"github.com/gin-gonic/gin"
)

func GinSetup() {
	router := gin.Default()

	initializeRoutes(router)

}

func initializeRoutes(router *gin.Engine) {
	router.Use(middlewares.MiddlewaresGet())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"sucess": "Up and running..."})
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Page not Found"})
	})
	router.GET("/all/:id", handlers.GetAllClient)
	router.GET("/consult", handlers.GetConsultClientID)
	router.POST("/insertclient", handlers.GetInsertCLientID)
	router.GET("/deleteclient/:id", handlers.DeleteClientId)
	router.POST("/updateclient", handlers.UpdateCleint)

	router.Run()
}
