package repositories

import (
	"log"
	"math/rand"
	"minisweeper/database"
	"minisweeper/domain"
)

type IGameRepository interface {
	Create(rows int, cols int, mines int) *domain.Game
	GetPoints(rows int, cols int) []domain.Point
	GetLast() *domain.Game
	AddMines(game *domain.Game) *domain.Game
}

type GameRepository struct {
	game *domain.Game
	DbConnection database.DbConnection
}

func NewGameRepository(connection database.DbConnection) IGameRepository {
	return &GameRepository{
		DbConnection: connection,
	}
}

func (r *GameRepository) Create(rows int, cols int, mines int) *domain.Game {
	queryBuilder := r.DbConnection.Connect()

	game := domain.NewGame(rows, cols, mines)
	game.AddPoints(r.GetPoints(rows, cols))

	result := queryBuilder.Create(&game) // pass pointer of data to Create

	if result.Error != nil {
		log.Println("some error", result.Error.Error())
	}

	r.game = game

	return game
}

func (r GameRepository) GetPoints(rows int, cols int) []domain.Point {
	var points []domain.Point

	for row := 1; row <= rows; row++ {
		for col := 1; col <= cols; col++ {
			point := domain.Point{
				Row: row,
				Col: col,
				Mine: false,
				Flag: false,
				Open: false,
				Value: 0,
			}

			points = append(points, point)
		}
	}

	return points
}

func (r GameRepository) AddMines(game *domain.Game) *domain.Game {
	queryBuilder := r.DbConnection.Connect()

	addedMines 	:= 0
	points 		:= len(game.Points)
	totalMines	:= game.Mines

	for  {
		key := rand.Intn(points)

		err := game.Points[key].AddMine()

		if err != nil {
			continue
		}

		queryBuilder.Save(&game.Points[key])
		addedMines++

		if addedMines == totalMines {
			break
		}
	}

	return game
}

func (r *GameRepository) GetLast() *domain.Game {
	queryBuilder := r.DbConnection.Connect()

	var game domain.Game

	queryBuilder.Preload("Points").Last(&game)

	return &game
}