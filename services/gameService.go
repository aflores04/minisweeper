package services

import (
	"minisweeper/game"
)

type GameService struct {

}

func NewGameService() *GameService {
	return &GameService{}
}

func (g GameService) Start(rows int, cols int, mines int) *game.Game {
	game := game.New(rows, cols, mines)

	game.SetPoints()
	game.SetMines()

	return game
}

