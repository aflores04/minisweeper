package game

type Point struct {
	X int		`json:"x"`
	Y int		`json:"y"`
	Mine bool	`json:"mine"`
	Flag bool	`json:"flag"`
	Value int 	`json:"value"`
	Open bool	`json:"open"`
}

func NewPoint(x int, y int, mine bool, value int) Point {
	return Point{
		X: x,
		Y: y,
		Mine: mine,
		Flag: false,
		Value: value,
		Open: false,
	}
}