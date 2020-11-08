package services

import (
	"minisweeper/game"
)

type GameService struct {

}

func NewGameService() *GameService {
	return &GameService{}
}

func (g *GameService) Start(rows int, cols int, mines int) *game.Game {
	newGame := game.New(rows, cols, mines)

	newGame.SetPoints()
	newGame.SetMines()

	return newGame
}

