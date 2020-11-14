package repositories_test

import (
	"github.com/go-playground/assert/v2"
	"minisweeper/database"
	"minisweeper/repositories"
	"testing"
)

var (
	connection = database.NewConnectionTest()
)

func getPointRepository() repositories.IPointRepository {
	return repositories.NewPointRepository(connection)
}

func TestPointRepository_AddRemoveFlag(t *testing.T) {
	pointRepository := getPointRepository()
	gameRepository := repositories.NewGameRepository(connection)

	game := gameRepository.Create(1,1,1)

	point := pointRepository.AddRemoveFlag(game.Points[0].ID, true)
	assert.Equal(t, true, point.Flag)

	pointWithOutFlag := pointRepository.AddRemoveFlag(game.Points[0].ID, false)
	assert.Equal(t, false, pointWithOutFlag.Flag)

	assert.PanicMatches(t, func() {
		pointRepository.AddRemoveFlag(999999, true)
	}, "point not found")
}