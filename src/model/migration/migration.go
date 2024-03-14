package migration

import (
	"github.com/google/uuid"
	"github.com/nazudis/disbursement/src/database"
	"github.com/nazudis/disbursement/src/model/entity"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func init() {
	db := database.GetDB()

	err := db.AutoMigrate(entity.Client{}, entity.Wallet{}, entity.Transaction{}, entity.AuthSession{})
	if err != nil {
		panic(err)
	}

	var client *entity.Client
	cid := uuid.MustParse("2606ed80-bbb9-4423-9867-eed4d7c5a3b3")
	secret := uuid.MustParse("be2a81e5-7dd2-4ea8-bf1a-f96755977bee")
	err = db.Model(&entity.Client{}).First(&client, "id = ?", cid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			client = &entity.Client{
				BaseEntityUID: entity.BaseEntityUID{ID: cid},
				Name:          "Default Client",
				Secret:        secret.String(),
			}
			err = db.Model(&entity.Client{}).Create(&client).Error
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	defaultAmount := decimal.NewFromInt(100000000)
	err = db.Model(&entity.Wallet{}).Where("owner_id = ?", client.ID).Update("amount", defaultAmount).Error
	if err != nil {
		panic(err)
	}

}
