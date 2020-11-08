package game

type Point struct {
	X int		`json:"x"`
	Y int		`json:"y"`
	Mine bool	`json:"mine"`
}

func NewPoint(x int, y int, mine bool) Point {
	return Point{
		X: x,
		Y: y,
		Mine: mine,
	}
}