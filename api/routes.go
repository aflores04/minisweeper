package api

import (
	"github.com/gin-gonic/gin"
	"minisweeper/database"
	"minisweeper/http/handlers"
	"minisweeper/repositories"
	"minisweeper/services"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	connection := database.NewConnection()
	gameRepository 	:= repositories.NewGameRepository(connection)
	gameService 	:= services.NewGameService(gameRepository)
	gameHandler 	:= handlers.NewGameHandler(gameService)

	v1 := router.Group("/api/v1")

	game := v1.Group("game")
	game.POST("/", gameHandler.CreateGameHandler)
	game.GET("/", gameHandler.CurrentGameHandler)

	point := game.Group("point")
	point.PUT("flag", gameHandler.FlagHandler)
	point.PUT("open", gameHandler.OpenPointHandler)

	return router
}