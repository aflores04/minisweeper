package services_test

import (
	"github.com/go-playground/assert/v2"
	"minisweeper/http/response"
	"minisweeper/repositories"
	"minisweeper/services"
	"testing"
)

func TestGameService_RemoveFlag(t *testing.T) {
	var flagResponse response.FlagResponse

	repository := repositories.NewGameRepository()

	service := services.NewGameService(repository)

	service.Start(2,2,1)

	flagResponse = service.AddRemoveFlag(1,1,false)

	assert.Equal(t, false, flagResponse.Flag)
}

func TestGameService_AddFlag(t *testing.T) {
	var flagResponse response.FlagResponse

	repository := repositories.NewGameRepository()

	service := services.NewGameService(repository)

	service.Start(2,2,1)

	flagResponse = service.AddRemoveFlag(1,1,true)

	assert.Equal(t, true, flagResponse.Flag)
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