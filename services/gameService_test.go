package services_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"minisweeper/database"
	"minisweeper/http/response"
	"minisweeper/repositories"
	"minisweeper/services"
	"testing"
)

func getService() services.IGameService {
	_ = godotenv.Load("../.env.local")

	connection := database.NewConnection()

	repository := repositories.NewGameRepository(connection)

	service := services.NewGameService(repository)

	return service
}

func TestGameService_Start(t *testing.T) {
	service := getService()
	game := service.Start(1,1,1)

	expectedGame := service.GetLast()

	assert.Equal(t, expectedGame, game)
}

func TestGameService_RemoveFlag(t *testing.T) {
	var pointResponse response.PointResponse

	service := getService()

	service.Start(2,2,1)

	pointResponse = service.AddRemoveFlag(1,1,false)

	assert.Equal(t, false, pointResponse.Flag)
}

func TestGameService_AddFlag(t *testing.T) {
	var pointResponse response.PointResponse

	service := getService()

	service.Start(2,2,1)

	pointResponse = service.AddRemoveFlag(1,1,true)

	assert.Equal(t, true, pointResponse.Flag)
}

func TestGameService_AddRemoveFlagWithNoGame(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, r, "there is no game running")
		}
	}()

	service := getService()
	service.AddRemoveFlag(1,1,true)
}

func TestGameService_OpenPointWithNoGame(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, r, "there is no game running")
		}
	}()

	service := getService()
	service.OpenPoint(1,1)
}

func TestGameService_OpenPointWithBadRequest(t *testing.T) {
	service := getService()
	service.Start(3,3,3)

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, r, "error in request")
		}
	}()

	service.OpenPoint(123,-123)
}

func TestGameService_OpenPoint(t *testing.T) {
	service := getService()
	service.Start(3,3,3)

	successResponse := service.OpenPoint(1,1)
	assert.Equal(t, true, successResponse.Open)
}