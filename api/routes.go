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
	pointRepository := repositories.NewPointRepository(connection)

	gameService 	:= services.NewGameService(gameRepository)
	pointService	:= services.NewPointService(pointRepository)

	gameHandler 	:= handlers.NewGameHandler(gameService)
	pointHandler	:= handlers.NewPointHandler(pointService)

	v1 := router.Group("/api/v1")

	game := v1.Group("game")
	game.POST("/", gameHandler.CreateGameHandler)
	game.GET("/", gameHandler.CurrentGameHandler)

	point := game.Group("point")
	point.PUT("flag", pointHandler.FlagHandler)
	point.PUT("open", gameHandler.OpenPointHandler)

	return router
}