package model

type ResponseSuccess struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type ResponseError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}
