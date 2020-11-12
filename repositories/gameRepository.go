package repositories

import (
	"minisweeper/game"
)

type IGameRepository interface {
	CreateGame(rows int, cols int, mines int) game.IGame
	GetGame() game.IGame
	AddRemoveFlag(row int, col int, flag bool) game.Point
	OpenPoint(row int, col int) game.Point
}

type GameRepository struct {
	game *game.Game
}

func NewGameRepository() IGameRepository {
	return &GameRepository{}
}

func (r *GameRepository) CreateGame(rows int, cols int, mines int) game.IGame {
	newGame := game.New(rows, cols, mines)
	newGame.SetMines()
	newGame.SetValues()

	r.game = newGame

	return newGame
}

func (r *GameRepository) GetGame() game.IGame {
	if r.game == nil {
		panic("there is no game running")
	}

	return r.game
}

func (r *GameRepository) AddRemoveFlag(row int, col int, flag bool) game.Point {
	if r.game != nil {
		for keyRow, rowPoint := range r.game.GetSquare().PointRows {
			for keyCol, point := range rowPoint.Points {
				if point.Y == row && point.X == col {
					r.game.Square.PointRows[keyRow].Points[keyCol].Flag = flag

					return *r.game.Square.PointRows[keyRow].Points[keyCol]
				}
			}
		}

		panic("error in request")
	}

	panic("there is no game running")
}

func (r *GameRepository) OpenPoint(row int, col int) game.Point {
	if r.game != nil {

		if row > r.game.Rows || col > r.game.Cols {
			panic("error in request")
		}
		for keyRow, rowPoint := range r.game.GetSquare().PointRows {
			for keyCol, point := range rowPoint.Points {
				if point.Y == row && point.X == col {
					r.game.Square.PointRows[keyRow].Points[keyCol].Open = true

					return *r.game.Square.PointRows[keyRow].Points[keyCol]
				}
			}
		}

	}

	panic("there is no game running")
}