package domain

type IGame interface {
	SetStatus(status string) IGame
	AddPoints(points []Point) IGame
}

type Game struct {
	ID 		uint 	`gorm:"primaryKey" json:"id"`
	Status	string	`json:"status"`
	Rows 	int		`json:"rows"`
	Cols	int		`json:"cols"`
	Mines	int 	`json:"mines"`
	Points	[]Point `gorm:"foreignKey:Game" json:"points"`
}

func NewGame(rows int, cols int, mines int) *Game {
	return &Game{
		Rows: rows,
		Cols: cols,
		Mines: mines,
		Status: "active",
	}
}

func (Game) TableName() string {
	return "games"
}

func (g *Game) SetStatus(status string) IGame {
	g.Status = status
	return g
}

func (g *Game) AddPoints(points []Point) IGame {
	g.Points = points
	return g
}