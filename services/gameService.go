package services

import (
	"minisweeper/game"
	"minisweeper/http/response"
	"minisweeper/repositories"
)

type IGameService interface {
	Start(rows int, cols int, mines int) game.IGame
	GetCurrent() game.IGame
	AddRemoveFlag(row int, col int, flag bool) response.FlagResponse
	GetPoint(row int, col int) game.Point
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

func (g *GameService) GetPoint(row int, col int) game.Point {
	return g.repository.GetPoint(row, col)
}

func (g *GameService) AddRemoveFlag(row int, col int, flag bool) response.FlagResponse {
	point := g.repository.AddRemoveFlag(row, col, flag)

	return response.FlagResponse{
		Row: point.Y,
		Col: point.X,
		Flag: point.Flag,
		Mine: point.Mine,
	}
}

