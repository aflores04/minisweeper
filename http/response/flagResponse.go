package response

import "minisweeper/domain"

type PointResponse struct {
	Code int `json:"code"`
	Point *domain.Point `json:"point"`
}
