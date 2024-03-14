package repository

import (
	"github.com/nazudis/disbursement/src/model/entity"
	"gorm.io/gorm"
)

type WalletRepositoryImpl struct {
	db *gorm.DB
}

// FirstByOwnerID implements WalletRepository.
func (w *WalletRepositoryImpl) FirstByOwnerID(id string) (*entity.Wallet, error) {
	var wallet entity.Wallet
	err := w.db.Where("owner_id = ?", id).First(&wallet).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &wallet, nil
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &WalletRepositoryImpl{db: db}
}
