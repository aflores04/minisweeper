package game

import (
	"errors"
	"math/rand"
)

type IGame interface {
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

	square := gameMap(rows, cols)

	return &Game{
		Cols: cols,
		Rows: rows,
		Mines: mines,
		Square: square,
	}
}

func gameMap(rows int, cols int) Square {
	var square Square

	for i := 0; i < rows; i++ {
		var pointRow PointRow
		for x := 0; x < cols; x++ {
			point := Point{
				X: x,
				Y: i,
				Mine: false,
				Value: 0,
				Flag: false,
			}
			pointRow.Points = append(pointRow.Points, &point)
		}
		square.PointRows = append(square.PointRows, &pointRow)
	}

	return square
}

func (g *Game) GetSquare() *Square {
	return &g.Square
}

func (g *Game) SetMines() {
	for i := 1; i <= g.Mines; i++ {
		randRow := rand.Intn(g.Rows)
		randCol := rand.Intn(g.Cols)

		if g.Square.PointRows[randRow].Points[randCol].Mine {
			g.SetMines()
		}

		if exceedTotalMines(g) {
			break
		}

		g.Square.PointRows[randRow].Points[randCol].Mine = true
	}
}

func exceedTotalMines(g *Game) bool {
	count := 0
	for _, pointRow := range g.Square.PointRows {
		for _, point := range pointRow.Points {
			if point.Mine {
				count++
			}
		}
	}

	if count >= g.Mines {
		return true
	}

	return false
}

func (g *Game) SetValues()  {
	for keyRow, pointRow := range g.Square.PointRows {
		for keyCol, point := range pointRow.Points {
			if point.Mine {

				if keyCol+1 <= len(pointRow.Points)-1 {
					g.PointRows[keyRow].Points[keyCol+1].Value++

					if keyRow+1 <= len(g.Square.PointRows)-1 {
						g.PointRows[keyRow+1].Points[keyCol+1].Value++
					}

					if keyRow-1 >= 0 {
						g.PointRows[keyRow-1].Points[keyCol+1].Value++
					}
				}

				if keyCol-1 >= 0 {
					g.PointRows[keyRow].Points[keyCol-1].Value++

					if keyRow+1 <= len(g.Square.PointRows)-1 {
						g.PointRows[keyRow+1].Points[keyCol-1].Value++
					}

					if keyRow-1 >= 0 {
						g.PointRows[keyRow-1].Points[keyCol-1].Value++
					}
				}

				if keyRow-1 >= 0 {
					g.PointRows[keyRow-1].Points[keyCol].Value++
				}

				if keyRow+1 <= len(g.Square.PointRows)-1 {
					g.PointRows[keyRow+1].Points[keyCol].Value++
				}

			}
		}
	}
}

func (g *Game) findNearlyMines(point Point) {
	//if nextPoint, err := g.FindPointByPosition(point.Y+1, point.X); err != nil {
	//	if nextPoint.Mine {
	//		g.Square.Points[nextPoint.Pos].Value++
	//	}
	//}
}

func (g *Game) FindPointByPosition(y int, x int) (*Point, error) {
	for _, pointRow := range g.Square.PointRows {
		for _, point := range pointRow.Points {
			if point.Y == y && point.X == x {
				return point, nil
			}
		}
	}

	return &Point{}, errors.New("can't find a point")
}