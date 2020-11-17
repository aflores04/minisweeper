package repositories

import (
	"minisweeper/database"
	"minisweeper/domain"
)

type IPointRepository interface {
	AddRemoveFlag(id uint, flag bool) *domain.Point
	Find(id uint) *domain.Point
	SumValue(point *domain.Point) *domain.Point
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
		panic("error in request")
	}

	point.Flag = flag

	queryBuilder.Save(&point)

	return &point
}

func (r PointRepository) Find (id uint) *domain.Point  {
	queryBuilder := r.DbConnection.Connect()

	var point domain.Point

	err := queryBuilder.First(&point, id)

	if err.Error != nil {
		panic("error in request")
	}

	return &point
}

func (r PointRepository) SumValue(point *domain.Point) *domain.Point {
	queryBuilder := r.DbConnection.Connect()

	point.Value = point.Value+1

	queryBuilder.Save(point)

	return point
}
