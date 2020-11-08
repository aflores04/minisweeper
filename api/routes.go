package api

import (
	"github.com/gin-gonic/gin"	
	"minisweeper/http/handlers"
	"minisweeper/services"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	gameService := services.NewGameService()
	gameHandler := handlers.NewGameHandler(gameService)

	router.POST("/game", gameHandler.CreateGameHandler)

	return router
}