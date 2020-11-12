package request

type PointRequest struct {
	Col int `json:"col"`
	Row int `json:"row"`
	Flag bool `json:"flag"`
	Open bool `json:"open"`
}
