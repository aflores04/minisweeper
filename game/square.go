package game

type Square struct {
	PointRows []*PointRow
}

type PointRow struct {
	Points []*Point
}

func NewSquare() *Square {
	return &Square{}
}