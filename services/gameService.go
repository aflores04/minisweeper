package services

import (
	"minisweeper/domain"
	"minisweeper/http/response"
	"minisweeper/repositories"
)

type IGameService interface {
	Start(rows int, cols int, mines int) *domain.Game
	GetLast() *domain.Game
	AddRemoveFlag(row int, col int, flag bool) response.PointResponse
	OpenPoint(row int, col int) response.PointResponse
}

type GameService struct {
	repository repositories.IGameRepository
}

func NewGameService(repository repositories.IGameRepository) IGameService {
	return &GameService{
		repository: repository,
	}
}

func (g *GameService) Start(rows int, cols int, mines int) *domain.Game {
	newGame := g.repository.Create(rows, cols, mines)
	g.repository.AddMines(newGame)

	return newGame
}

func (g *GameService) GetLast() *domain.Game {
	return g.repository.GetLast()
}

func (g *GameService) AddRemoveFlag(row int, col int, flag bool) response.PointResponse {
	//point := g.repository.AddRemoveFlag(row, col, flag)
	//
	//return response.PointResponse{
	//	Row: point.Y,
	//	Col: point.X,
	//	Flag: point.Flag,
	//	Value: point.Value,
	//	Mine: point.Mine,
	//}
	return response.PointResponse{}
}

func (g *GameService) OpenPoint(row int, col int) response.PointResponse {
	//point := g.repository.OpenPoint(row, col)
	//
	//return response.PointResponse{
	//	Row: point.Y,
	//	Col: point.X,
	//	Flag: point.Flag,
	//	Mine: point.Mine,
	//	Value: point.Value,
	//	Open: point.Open,
	//}
	return response.PointResponse{}
}