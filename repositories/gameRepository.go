package repositories

import (
	"log"
	"minisweeper/database"
	"minisweeper/domain"
)

type IGameRepository interface {
	Create(rows int, cols int, mines int) *domain.Game
	GetPoints(rows int, cols int) []domain.Point
	GetLast() *domain.Game
	//AddRemoveFlag(row int, col int, flag bool) game.Point
	//OpenPoint(row int, col int) game.Point
}

type GameRepository struct {
	game *domain.Game
	DbConnection database.DbConnection
}

func NewGameRepository(connection database.DbConnection) IGameRepository {
	return &GameRepository{
		DbConnection: connection,
	}
}

func (r *GameRepository) Create(rows int, cols int, mines int) *domain.Game {
	game := domain.NewGame(rows, cols, mines)
	game.AddPoints(r.GetPoints(rows, cols))

	queryBuilder := r.DbConnection.Connect()

	result := queryBuilder.Create(&game) // pass pointer of data to Create

	if result.Error != nil {
		log.Println("some error", result.Error.Error())
	}

	return game
}

func (r *GameRepository) GetPoints(rows int, cols int) []domain.Point {
	var points []domain.Point

	for row := 1; row <= rows; row++ {
		for col := 1; col <= cols; col++ {
			point := domain.Point{
				Row: row,
				Col: col,
				Mine: false,
				Flag: false,
				Open: false,
				Value: 0,
			}

			points = append(points, point)
		}
	}

	return points
}

func (r *GameRepository) GetLast() *domain.Game {
	queryBuilder := r.DbConnection.Connect()

	var game domain.Game

	queryBuilder.Preload("Points").Last(&game)

	return &game
}

//func (r *GameRepository) AddRemoveFlag(row int, col int, flag bool) game.Point {
//	if r.game != nil {
//		for keyRow, rowPoint := range r.game.GetSquare().PointRows {
//			for keyCol, point := range rowPoint.Points {
//				if point.Y == row && point.X == col {
//					r.game.Square.PointRows[keyRow].Points[keyCol].Flag = flag
//
//					return *r.game.Square.PointRows[keyRow].Points[keyCol]
//				}
//			}
//		}
//
//		panic("error in request")
//	}
//
//	panic("there is no game running")
//}

//func (r *GameRepository) OpenPoint(row int, col int) game.Point {
//	if r.game != nil {
//
//		if row > r.game.Rows || col > r.game.Cols {
//			panic("error in request")
//		}
//		for keyRow, rowPoint := range r.game.GetSquare().PointRows {
//			for keyCol, point := range rowPoint.Points {
//				if point.Y == row && point.X == col {
//					r.game.Square.PointRows[keyRow].Points[keyCol].Open = true
//
//					return *r.game.Square.PointRows[keyRow].Points[keyCol]
//				}
//			}
//		}
//
//	}
//
//	panic("there is no game running")
//}