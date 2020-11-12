package domain

type Point struct {
	ID		uint `gorm:"primaryKey" json:"-"`
	Row		int	`json:"row"`
	Col		int `json:"col"`
	Mine	bool `json:"mine"`
	Flag	bool `json:"flag"`
	Value 	int `json:"value"`
	Open 	bool `json:"open"`
	Game 	uint `json:"-"`
}