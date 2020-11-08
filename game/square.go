package game

type Square struct {
	Points []Point `json:"points"`
}

func NewSquare() *Square {
	return &Square{}
}