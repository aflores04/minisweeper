package services

import (
	"minisweeper/domain"
	"minisweeper/repositories"
)

type IPointService interface {
	AddRemoveFlag(id uint, flag bool) *domain.Point
}

type PointService struct {
	repository repositories.IPointRepository
}

func NewPointService(repository repositories.IPointRepository) IPointService {
	return &PointService{
		repository: repository,
	}
}

func (s PointService) AddRemoveFlag(id uint, flag bool) *domain.Point {
	point := s.repository.AddRemoveFlag(id, flag)

	return point
}