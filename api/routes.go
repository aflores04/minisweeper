package api

import (
	"github.com/gin-gonic/gin"	
	"minisweeper/http/handlers"
	"minisweeper/repositories"
	"minisweeper/services"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	gameRepository 	:= repositories.NewGameRepository()
	gameService 	:= services.NewGameService(gameRepository)
	gameHandler 	:= handlers.NewGameHandler(gameService)

	router.POST("/game", gameHandler.CreateGameHandler)
	router.PUT("/game/point/flag", gameHandler.FlagHandler)
	router.PUT("/game/point/open", gameHandler.OpenPointHandler)
	router.GET("/game/current", gameHandler.CurrentGameHandler)

	return router
}