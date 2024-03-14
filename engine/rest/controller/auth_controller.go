package controller

import "net/http"

type AuthController interface {
	GetAccessToken(w http.ResponseWriter, r *http.Request)
}
