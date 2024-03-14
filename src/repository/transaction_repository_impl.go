package repository

import (
	"fmt"

	"github.com/nazudis/disbursement/src/model/entity"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

// FirstByReferenceID implements TransactionRepository.
func (r *TransactionRepositoryImpl) FirstByReferenceID(refId string) (*entity.Transaction, error) {
	var trx entity.Transaction
	err := r.db.Where("reference_id = ?", refId).First(&trx).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &trx, nil
}

// TrxFailed implements TransactionRepository.
func (r *TransactionRepositoryImpl) TrxFailed(refId string) (*entity.Transaction, error) {
	trx, err := r.FirstByReferenceID(refId)
	if err != nil {
		return nil, err
	}

	if trx == nil {
		return nil, fmt.Errorf("transaction not found")
	}

	trx.Status = entity.TrxStatusFailed
	err = r.db.Save(trx).Error
	if err != nil {
		return nil, err
	}

	var wallet *entity.Wallet
	r.db.Where("owner_id = ?", trx.OwnerID).First(&wallet)
	wallet.Amount = wallet.Amount.Add(trx.Amount)
	err = r.db.Save(wallet).Error
	if err != nil {
		return nil, err
	}

	return trx, nil
}

// TrxCompleted implements TransactionRepository.
func (r *TransactionRepositoryImpl) TrxCompleted(refId string) (*entity.Transaction, error) {
	trx, err := r.FirstByReferenceID(refId)
	if err != nil {
		return nil, err
	}

	if trx == nil {
		return nil, fmt.Errorf("transaction not found")
	}

	trx.Status = entity.TrxStatusCompleted
	err = r.db.Save(trx).Error
	if err != nil {
		return nil, err
	}

	return trx, nil
}

// Disbursement implements TransactionRepository.
func (r *TransactionRepositoryImpl) Disbursement(trx *entity.Transaction, wallet *entity.Wallet) error {
	err := r.db.Create(trx).Error
	if err != nil {
		return err
	}

	wallet.Amount = wallet.Amount.Sub(trx.Amount)
	err = r.db.Save(wallet).Error
	if err != nil {
		return err
	}

	return nil
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}
