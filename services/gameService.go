package services

import (
	"minisweeper/game"
	"minisweeper/repositories"
)

type GameService struct {
	repository *repositories.GameRepository
}

func NewGameService(repository *repositories.GameRepository) *GameService {
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

