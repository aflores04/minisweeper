package response

import "minisweeper/domain"

type CreateGameResponse struct {
	Code int `json:"code"`
	Game *domain.Game `json:"game"`
}