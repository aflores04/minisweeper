package request

type FlagRequest struct {
	Col int `json:"col"`
	Row int `json:"row"`
	Flag bool `json:"flag"`
}
