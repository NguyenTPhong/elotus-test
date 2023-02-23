package entity

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Error *ResponseError `json:"error,omitempty"`
	Data  interface{}    `json:"data,omitempty"`
}
