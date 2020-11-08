package handlers

import (
	"minisweeper/services"
)

type GameHandler struct {
	Service services.IGameService
}

func NewGameHandler(service services.IGameService) *GameHandler {
	return &GameHandler{
		Service: service,
	}
}




