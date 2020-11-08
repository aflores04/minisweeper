package request

type AddFlagRequest struct {
	Col int `json:"col"`
	Row int `json:"row"`
	Flag bool `json:"flag"`
}
