package game

import (
	"math/rand"
)

type IGame interface {
	SetPoints()
	SetMines()
	GetSquare() *Square
}

type Game struct {
	Cols int `json:"cols"`
	Rows int `json:"rows"`
	Mines int `json:"mines"`
	Square `json:"square"`
}

func New(cols int, rows int, mines int) *Game {

	if cols <= 0 || rows <= 0 || mines <= 0 {
		panic("error in request")
	}

	return &Game{
		Cols: cols,
		Rows: rows,
		Mines: mines,
	}
}

func (g *Game) GetSquare() *Square {
	return &g.Square
}

func (g *Game) SetPoints() {	
	points := []Point{}

	for y := 1; y <= g.Rows; y++ {
		for x := 1; x <= g.Cols; x++ {
			point := NewPoint(y, x, false)

			points = append(points, point)
		}
	}

	g.Square.Points = points
}

func (g *Game) SetMines() {
	for i := 1; i <= g.Mines; i++ {
		randPoint := rand.Intn(len(g.Square.Points))
		
		g.Square.Points[randPoint].Mine = true
	}
}
