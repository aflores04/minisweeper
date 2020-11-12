package services_test

import (
	"github.com/go-playground/assert/v2"
	"minisweeper/http/response"
	"minisweeper/repositories"
	"minisweeper/services"
	"testing"
)

func TestGameService_RemoveFlag(t *testing.T) {
	var pointResponse response.PointResponse

	repository := repositories.NewGameRepository()

	service := services.NewGameService(repository)

	service.Start(2,2,1)

	pointResponse = service.AddRemoveFlag(1,1,false)

	assert.Equal(t, false, pointResponse.Flag)
}

func TestGameService_AddFlag(t *testing.T) {
	var pointResponse response.PointResponse

	repository := repositories.NewGameRepository()

	service := services.NewGameService(repository)

	service.Start(2,2,1)

	pointResponse = service.AddRemoveFlag(1,1,true)

	assert.Equal(t, true, pointResponse.Flag)
}

func TestGameService_AddRemoveFlagWithNoGame(t *testing.T) {
	repository := repositories.NewGameRepository()

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, r, "there is no game running")
		}
	}()

	service := services.NewGameService(repository)

	service.AddRemoveFlag(1,1,true)
}

func TestGameService_OpenPointWithNoGame(t *testing.T) {
	repository := repositories.NewGameRepository()

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, r, "there is no game running")
		}
	}()

	service := services.NewGameService(repository)

	service.OpenPoint(1,1)
}

func TestGameService_OpenPointWithBadRequest(t *testing.T) {
	repository := repositories.NewGameRepository()

	service := services.NewGameService(repository)
	service.Start(3,3,3)

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, r, "error in request")
		}
	}()

	service.OpenPoint(123,-123)
}

func TestGameService_OpenPoint(t *testing.T) {
	repository := repositories.NewGameRepository()

	service := services.NewGameService(repository)
	service.Start(3,3,3)

	successResponse := service.OpenPoint(1,1)

	assert.Equal(t, true, successResponse.Open)
}