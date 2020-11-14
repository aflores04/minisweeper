package request

type PointRequest struct {
	ID uint `json:"id"`
	Flag bool `json:"flag"`
	Open bool `json:"open"`
}
