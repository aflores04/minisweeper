package response

type PointResponse struct {
	Code int `json:"code"`
	Row int	`json:"row"`
	Col int `json:"col"`
	Value int `json:"value"`
	Flag bool `json:"flag"`
	Mine bool `json:"mine"`
	Open bool `json:"open"`
}
