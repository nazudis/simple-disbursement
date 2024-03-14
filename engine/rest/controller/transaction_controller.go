package controller

import "net/http"

type TransactionController interface {
	GetAccountVerification(w http.ResponseWriter, r *http.Request)
	PostDisbursement(w http.ResponseWriter, r *http.Request)
	PostCallbackDisbursement(w http.ResponseWriter, r *http.Request)
}
