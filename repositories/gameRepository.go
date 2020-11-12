package repositories

import (
	"log"
	"minisweeper/database"
	"minisweeper/domain"
	"minisweeper/game"
)

type IGameRepository interface {
	Create(rows int, cols int, mines int) domain.Game
	GetLast() domain.Game
	AddRemoveFlag(row int, col int, flag bool) game.Point
	OpenPoint(row int, col int) game.Point
}

type GameRepository struct {
	game *game.Game
	DbConnection database.DbConnection
}

func NewGameRepository(connection database.DbConnection) IGameRepository {
	return &GameRepository{
		DbConnection: connection,
	}
}

func (r *GameRepository) Create(rows int, cols int, mines int) domain.Game {
	game := domain.Game{
		Rows: rows,
		Cols: cols,
		Mines: mines,
		Status: "active",
	}

	queryBuilder := r.DbConnection.Connect()

	result := queryBuilder.Create(&game) // pass pointer of data to Create

	setPoints(&game)

	if result.Error != nil {
		log.Println("some error", result.Error.Error())
	}

	return game
}

func setPoints(game *domain.Game) {

	//for row := 0; row <= game.Rows; i++ {
	//	for col := 0; col <= game.Cols; i++ {
	//		point := domain.Point{
	//			Row:
	//		}
	//	}
	//}

}

func (r *GameRepository) GetLast() domain.Game {
	queryBuilder := r.DbConnection.Connect()

	var game domain.Game

	queryBuilder.Last(&game)

	return game
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