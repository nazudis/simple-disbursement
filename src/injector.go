//go:build wireinject
// +build wireinject

package src

import (
	"github.com/google/wire"
	"github.com/nazudis/disbursement/engine/rest/controller"
	"github.com/nazudis/disbursement/engine/rest/transformer"
	"github.com/nazudis/disbursement/src/database"
	"github.com/nazudis/disbursement/src/dependencies/mockapi"
	"github.com/nazudis/disbursement/src/repository"
	"github.com/nazudis/disbursement/src/service"
)

func InitializeAuthController() (controller.AuthController, error) {
	wire.Build(
		controller.NewAuthController,
		service.NewAuthService,
		transformer.NewAuthTransformer,
		repository.NewClientRepository,
		repository.NewAuthSessionRepository,
		database.GetDB,
	)

	return nil, nil
}

func InitializeTransactionController() (controller.TransactionController, error) {
	wire.Build(
		controller.NewTransactionController,
		service.NewTransactionService,
		transformer.NewTransactionTransformer,
		mockapi.NewMockApi,
		repository.NewTransactionRepository,
		repository.NewWalletRepository,
		database.GetDB,
	)

	return nil, nil
}
