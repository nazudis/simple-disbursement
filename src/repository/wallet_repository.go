package repository

import "github.com/nazudis/disbursement/src/model/entity"

type WalletRepository interface {
	FirstByOwnerID(id string) (*entity.Wallet, error)
}
