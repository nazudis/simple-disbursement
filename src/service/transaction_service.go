package service

import (
	mockapiResponse "github.com/nazudis/disbursement/src/dependencies/mockapi/response"
	"github.com/nazudis/disbursement/src/model/entity"
)

type TransactionService interface {
	GetAccountVerification(accountNumber, bankShortCode string) (*mockapiResponse.Account, error)
	PostDisbursement(params PostDisbursementParams) (*entity.Transaction, error)
	PostCallbackDisbursement(params PostCallbackDisbursementParams) error
}
