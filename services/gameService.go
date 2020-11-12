package services

import (
	"minisweeper/game"
	"minisweeper/http/response"
	"minisweeper/repositories"
)

type IGameService interface {
	Start(rows int, cols int, mines int) game.IGame
	GetCurrent() game.IGame
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

func (g *GameService) Start(rows int, cols int, mines int) game.IGame {
	newGame := g.repository.CreateGame(rows, cols, mines)

	return newGame
}

func (g *GameService) GetCurrent() game.IGame {
	return g.repository.GetGame()
}

func (g *GameService) AddRemoveFlag(row int, col int, flag bool) response.PointResponse {
	point := g.repository.AddRemoveFlag(row, col, flag)

	return response.PointResponse{
		Row: point.Y,
		Col: point.X,
		Flag: point.Flag,
		Value: point.Value,
		Mine: point.Mine,
	}
}

func (g *GameService) OpenPoint(row int, col int) response.PointResponse {
	point := g.repository.OpenPoint(row, col)

	return response.PointResponse{
		Row: point.Y,
		Col: point.X,
		Flag: point.Flag,
		Mine: point.Mine,
		Value: point.Value,
		Open: point.Open,
	}
}