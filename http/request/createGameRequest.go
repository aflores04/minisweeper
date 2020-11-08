package request

type CreateGameRequest struct {
	Cols 	int `json:"cols" binding:"required"`
	Rows 	int `json:"rows" binding:"required"`
	Mines 	int	`json:"mines" binding:"required"`
}