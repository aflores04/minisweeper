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
	repository 		repositories.IGameRepository
	pointRepository	repositories.IPointRepository
}

func NewGameService(repository repositories.IGameRepository, pointRepository repositories.IPointRepository) IGameService {
	return &GameService{
		repository: repository,
		pointRepository: pointRepository,
	}
}

func (g *GameService) Start(rows int, cols int, mines int) *domain.Game {
	newGame := g.repository.Create(rows, cols, mines)
	g.repository.AddMines(newGame)
	g.addValues(newGame)

	// need to return when values are assign to the points
	game := g.repository.Find(newGame.ID)

	return game
}

func (g GameService) addValues(game *domain.Game) {
	for row := 1; row <= game.Rows; row++ { // row 1
		for col := 1; col <= game.Cols; col++ { // col 2
			point, _ := g.repository.GetPointByPos(row, col) //
			if point.Mine { // true
				g.findAndSum(row-1, col-1) // 0 1
				g.findAndSum(row, col-1) // 1 1
				g.findAndSum(row+1, col-1) // 2 1

				g.findAndSum(row-1, col) // 0 2
				g.findAndSum(row+1, col) // 2 2

				g.findAndSum(row-1, col+1) // 0 3
				g.findAndSum(row, col+1) // 1 3
				g.findAndSum(row+1, col+1) // 2 3
			}
		}
	}
}

func (g GameService) findAndSum(row int, col int) {
	if point, err := g.repository.GetPointByPos(row, col); err == nil {
		g.pointRepository.SumValue(point)
	}
}

func (g *GameService) GetLast() *domain.Game {
	return g.repository.GetLast()
}