package services_test

import (
	"github.com/go-playground/assert/v2"
	"minisweeper/repositories"
	"minisweeper/services"
	"testing"
)

func TestGameService_RemoveFlag(t *testing.T) {
	repository := repositories.NewGameRepository()

	service := services.NewGameService(repository)

	service.Start(2,2,1)

	service.AddRemoveFlag(1,1,false)

	assert.Equal(t, false, service.GetPoint(1, 1).Flag)
}

func TestGameService_AddFlag(t *testing.T) {
	repository := repositories.NewGameRepository()

	service := services.NewGameService(repository)

	service.Start(2,2,1)

	service.AddRemoveFlag(1,1,true)

	assert.Equal(t, true, service.GetPoint(1, 1).Flag)
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