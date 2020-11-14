package repositories

import (
	"minisweeper/database"
	"minisweeper/domain"
)

type IPointRepository interface {
	AddRemoveFlag(id uint, flag bool) *domain.Point
}

type PointRepository struct {
	point *domain.Point
	DbConnection database.DbConnection
}

func NewPointRepository(connection database.DbConnection) IPointRepository {
	return &PointRepository{
		DbConnection: connection,
	}
}

func (r *PointRepository) AddRemoveFlag(id uint, flag bool) *domain.Point {
	queryBuilder := r.DbConnection.Connect()

	var point domain.Point

	err := queryBuilder.First(&point, id)

	if err.Error != nil {
		panic("point not found")
	}

	point.Flag = flag

	queryBuilder.Save(&point)

	return &point
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