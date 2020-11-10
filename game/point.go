package game

type Point struct {
	Pos int		`json:"-"`
	X int		`json:"x"`
	Y int		`json:"y"`
	Mine bool	`json:"mine"`
	Flag bool	`json:"flag"`
	Value int 	`json:"value"`
}

func NewPoint(x int, y int, mine bool, value int) Point {
	return Point{
		X: x,
		Y: y,
		Mine: mine,
		Flag: false,
		Value: value,
	}
}