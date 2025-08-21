package utils


type Response struct {
	Status 	bool  		`json:"status"`
	Data  	interface{} `json:"data,omitempty"`
	Message string		`json:"message"`
}

func ResponseFull(status bool, data interface{}, message string) *Response{
	return &Response{
		Status: status,
		Data: data,
		Message: message,
	}
}

func ResponseNotData(status bool, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}