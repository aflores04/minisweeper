package repositories

import (
	"github.com/stretchr/testify/assert"
	"minisweeper/database"
	"minisweeper/domain"
	"testing"
)

func getRepository() IGameRepository {
	connection := database.NewConnectionTest()

	repository := NewGameRepository(connection)

	return repository
}

func TestGameRepository_CreateGame(t *testing.T) {
	repository := getRepository()
	newGame := repository.Create(2,2,2)

	expectedGame := repository.GetLast()

	assert.Equal(t, expectedGame, newGame)
}

func TestGameRepository_GetPoints(t *testing.T) {
	repository := getRepository()

	points := repository.GetPoints(2,2)

	expectedPoints := []domain.IPoint{
		domain.NewPoint(1,1,false,false,0,false),
		domain.NewPoint(1,2,false,false,0,false),
		domain.NewPoint(2,1,false,false,0,false),
		domain.NewPoint(2,2,false,false,0,false),
	}

	assert.Equal(t, expectedPoints, points)
}