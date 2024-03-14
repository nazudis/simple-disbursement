package repository

import (
	"github.com/nazudis/disbursement/src/model/entity"
	"gorm.io/gorm"
)

type ClientRepositoryImpl struct {
	db *gorm.DB
}

// FirstByID implements ClientRepository.
func (r *ClientRepositoryImpl) FirstByID(id string) (*entity.Client, error) {
	var client entity.Client
	if err := r.db.First(&client, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &client, nil
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &ClientRepositoryImpl{db: db}
}
