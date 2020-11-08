package handlers

import (
	"minisweeper/services"
)

type GameHandler struct {
	Service *services.GameService
}

func NewGameHandler(service *services.GameService) *GameHandler {
	return &GameHandler{
		Service: services.NewGameService(),
	}
}




