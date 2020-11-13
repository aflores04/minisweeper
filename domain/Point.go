package domain

import "errors"

type IPoint interface {
	GetRow() int
	GetCol() int
}

type Point struct {
	ID		uint `gorm:"primaryKey" json:"id"`
	Row		int	`json:"row"`
	Col		int `json:"col"`
	Mine	bool `json:"mine"`
	Flag	bool `json:"flag"`
	Value 	int `json:"value"`
	Open 	bool `json:"open"`
	Game 	uint `json:"-"`
}

func NewPoint(row int, col int, mine bool, flag bool, value int, open bool) *Point {
	return &Point{
		Row: row,
		Col: col,
		Mine: mine,
		Flag: flag,
		Value: value,
		Open: open,
	}
}

func (Point) TableName() string {
	return "points"
}

func (p Point) GetRow() int {
	return p.Row
}

func (p Point) GetCol() int {
	return p.Col
}

func (p *Point) AddMine() error {
	if p.Mine {
		return errors.New("already got a mine")
	}

	p.Mine = true

	return nil
}