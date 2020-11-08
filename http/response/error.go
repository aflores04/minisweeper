package response

type ErrorResponse struct {
	Code 	int		`json:"code"`
	Message	string	`json:"message"`	
}

func (ErrorResponse) Error() interface{} {
	return nil
}