package handlers

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponse(data interface{}, message string) *Response {
	return &Response{data, message}
}
