package repositories_test

import (
	"github.com/stretchr/testify/assert"
	"minisweeper/database"
	"minisweeper/domain"
	"minisweeper/repositories"
	"testing"
)

func getGameRepository() repositories.IGameRepository {
	connection := database.NewConnectionTest()

	repository := repositories.NewGameRepository(connection)

	return repository
}

func TestGameRepository_CreateGame(t *testing.T) {
	repository := getGameRepository()
	newGame := repository.Create(2,2,2)

	expectedGame := repository.GetLast()

	assert.Equal(t, expectedGame, newGame)
}

func TestGameRepository_GetPoints(t *testing.T) {
	repository := getGameRepository()

	points := repository.GetPoints(2,2)

	expectedPoints := []domain.IPoint{
		domain.NewPoint(1,1,false,false,0,false),
		domain.NewPoint(1,2,false,false,0,false),
		domain.NewPoint(2,1,false,false,0,false),
		domain.NewPoint(2,2,false,false,0,false),
	}

	assert.Equal(t, expectedPoints, points)
}

func TestGameRepository_AddMines(t *testing.T) {
	repository := getGameRepository()

	newGame := repository.Create(2,2,3)
	game := repository.AddMines(newGame)

	var mines int = 0

	for _, point := range game.Points {
		if point.Mine {
			mines++
		}
	}

	assert.Equal(t, 3, mines)
}