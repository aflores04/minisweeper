package repositories

import (
	"minisweeper/game"
)

type IGameRepository interface {
	CreateGame(rows int, cols int, mines int) game.IGame
	GetGame() game.IGame
	AddRemoveFlag(row int, col int, flag bool) game.Point
	GetPoint(row int, col int) game.Point
}

type GameRepository struct {
	game *game.Game
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

func (r *GameRepository) AddRemoveFlag(row int, col int, flag bool) game.Point {
	if r.game != nil {
		for key, point := range r.game.GetSquare().Points {
			if point.Y == row && point.X == col {
				r.game.Square.Points[key].Flag = flag

				return r.game.Square.Points[key]
			}
		}

		panic("error in request")
	}

	panic("there is no game running")
}

func (r *GameRepository) GetPoint(row int, col int) game.Point {
	if r.game != nil {
		for _, point := range r.game.GetSquare().Points {
			if point.Y == row && point.X == col {
				return point
			}
		}
	}

	panic("there is no game running")
}