package services_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"minisweeper/database"
	"minisweeper/repositories"
	"minisweeper/services"
	"testing"
)

func getService() services.IGameService {
	_ = godotenv.Load("../.env.local")

	connection := database.NewConnection()

	repository := repositories.NewGameRepository(connection)
	pointRepository := repositories.NewPointRepository(connection)

	service := services.NewGameService(repository, pointRepository)

	return service
}

func TestGameService_Start(t *testing.T) {
	service := getService()
	game := service.Start(4,4,3)

	expectedGame := service.GetLast()

	assert.Equal(t, expectedGame, game)
}