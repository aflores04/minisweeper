package domain

type Game struct {
	ID 		uint 	`gorm:"primaryKey" json:"id"`
	Status	string	`json:"status"`
	Rows 	int		`json:"rows"`
	Cols	int		`json:"cols"`
	Mines	int 	`json:"mines"`
	Points	[]Point `gorm:"foreignKey:Game" json:"points"`
}

