package repository

import "github.com/nazudis/disbursement/src/model/entity"

type TransactionRepository interface {
	FirstByReferenceID(refId string) (*entity.Transaction, error)
	Disbursement(trx *entity.Transaction, wallet *entity.Wallet) error
	TrxCompleted(refId string) (*entity.Transaction, error)
	TrxFailed(refId string) (*entity.Transaction, error)
}
