package repository

import "github.com/nazudis/disbursement/src/model/entity"

type ClientRepository interface {
	FirstByID(id string) (*entity.Client, error)
}
