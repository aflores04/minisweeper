package repositories

import (
	"log"
	"minisweeper/game"
)

type IGameRepository interface {
	CreateGame(rows int, cols int, mines int) game.IGame
	GetGame() game.IGame
	AddRemoveFlag(row int, col int, flag bool)
}

type GameRepository struct {
	game game.IGame
}

func NewGameRepository() IGameRepository {
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

func (r *GameRepository) AddRemoveFlag(row int, col int, flag bool) {
	for point := range r.game.GetSquare().Points {
		log.Println(point)
	}
}