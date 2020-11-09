package response

type FlagResponse struct {
	Code int `json:"code"`
	Row int	`json:"row"`
	Col int `json:"col"`
	Flag bool `json:"flag"`
	Mine bool `json:"mine"`
}
