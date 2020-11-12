package repositories

import (
	"github.com/joho/godotenv"
	"minisweeper/database"
	"testing"
)

func TestGameRepository_CreateGame(t *testing.T) {
	_ = godotenv.Load("../.env.local")

	connection := database.NewConnection()

	repository := NewGameRepository(connection)

	repository.CreateGame(1,1,1)
}
