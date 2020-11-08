package api

import (
	"github.com/gin-gonic/gin"	
	"minisweeper/handlers"
	"minisweeper/services"
)

func InitRoutes(router *gin.Engine) *gin.Engine {
	gameService := services.NewGameService()
	gameHandler := handlers.NewGameHandler(gameService)

	router.POST("/game", gameHandler.NewGameHandler)

	return router
}