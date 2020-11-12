package repositories

import (
	"minisweeper/database"
	"testing"
)

func TestGameRepository_CreateGame(t *testing.T) {
	connection := database.NewConnectionTest()

	repository := NewGameRepository(connection)

	repository.Create(1,1,1)
}
