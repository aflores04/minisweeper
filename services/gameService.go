package services

import (
	"minisweeper/domain"
	"minisweeper/repositories"
)

type IGameService interface {
	Start(rows int, cols int, mines int) *domain.Game
	GetLast() *domain.Game
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