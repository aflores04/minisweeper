package repositories

import (
	"minisweeper/game"
)

type GameRepository struct {
	game game.IGame
}

func NewGameRepository() *GameRepository {
	return &GameRepository{}
}

func (r *GameRepository) CreateGame(rows int, cols int, mines int) game.IGame {
	newGame := game.New(rows, cols, mines)
	newGame.SetPoints()
	newGame.SetMines()

	r.game = newGame

	return newGame
}

func (r *GameRepository) GetGame() game.IGame {
	return r.game
}