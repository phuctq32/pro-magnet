package common

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Count   *int        `json:"count,omitempty"`
}

func NewResponse(msg string, data interface{}) *Response {
	return &Response{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func NewArrayResponse(msg string, data interface{}, count int) *Response {
	return &Response{
		Success: true,
		Message: msg,
		Data:    data,
		Count:   &count,
	}
}
