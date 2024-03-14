package helper

import "net/http"

type StdResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ReplySuccess(message string, data any) StdResponse {
	resData := StdResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}
	return resData
}

func ReplyFail(code int, message string) StdResponse {
	resData := StdResponse{
		Status:  code,
		Message: message,
	}
	return resData
}
