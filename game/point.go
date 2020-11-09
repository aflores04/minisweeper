package game

type Point struct {
	ID int		`json:"id"`
	X int		`json:"x"`
	Y int		`json:"y"`
	Mine bool	`json:"mine"`
	Flag bool	`json:"flag"`
}

func NewPoint(x int, y int, mine bool) *Point {
	return &Point{
		X: x,
		Y: y,
		Mine: mine,
		Flag: false,
	}
}