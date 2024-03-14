package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nazudis/disbursement/src/dependencies/mockapi"
	mockapiRequest "github.com/nazudis/disbursement/src/dependencies/mockapi/request"
	mockapiResponse "github.com/nazudis/disbursement/src/dependencies/mockapi/response"
	"github.com/nazudis/disbursement/src/model/entity"
	"github.com/nazudis/disbursement/src/repository"
)

type TransactionServiceImpl struct {
	MockApi               mockapi.MockApi
	TransactionRepository repository.TransactionRepository
	WalletRepository      repository.WalletRepository
}

// GetAccountVerification implements TransactionService.
func (s *TransactionServiceImpl) GetAccountVerification(accountNumber string, bankShortCode string) (*mockapiResponse.Account, error) {
	acc, err := s.MockApi.GetAccountByAccountNumberAndBankShortCode(accountNumber, bankShortCode)
	if err != nil {
		return nil, fmt.Errorf("invalid account number. please check your account number and bank short code")
	}
	return acc, nil
}

// PostDisbursement implements TransactionService.
func (s *TransactionServiceImpl) PostDisbursement(params PostDisbursementParams) (*entity.Transaction, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}

	wallet, err := s.WalletRepository.FirstByOwnerID(params.OwnerID)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, fmt.Errorf("wallet not found")
	}

	if params.Amount.GreaterThan(wallet.Amount) {
		return nil, fmt.Errorf("the request cannot be process because insufficient balance")
	}

	exists, err := s.TransactionRepository.FirstByReferenceID(params.ReferenceID)
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return nil, fmt.Errorf("the reference id is already exists")
	}

	bankTrx, err := s.MockApi.PostTransaction(mockapiRequest.PostTransactionRequest{
		Type:          params.Type,
		Amount:        params.Amount.InexactFloat64(),
		AccountNumber: params.BankAccountNo,
	})
	if err != nil {
		return nil, err
	}

	trx := &entity.Transaction{
		OwnerID:                uuid.MustParse(params.OwnerID),
		ReferenceID:            params.ReferenceID,
		ExtReferenceID:         bankTrx.TrxID,
		Description:            params.Description,
		Type:                   params.Type,
		Amount:                 params.Amount,
		RecipientAccountNumber: params.BankAccountNo,
		Status:                 entity.TrxStatusProcessing,
	}

	err = s.TransactionRepository.Disbursement(trx, wallet)
	if err != nil {
		return nil, err
	}

	return trx, nil
}

// PostCallbackDisbursement implements TransactionService.
func (s *TransactionServiceImpl) PostCallbackDisbursement(params PostCallbackDisbursementParams) (err error) {
	switch params.Status {
	case entity.TrxStatusCompleted:
		_, err = s.TransactionRepository.TrxCompleted(params.ReferenceID)
	case entity.TrxStatusFailed:
		_, err = s.TransactionRepository.TrxFailed(params.ReferenceID)
	}
	return
}

func NewTransactionService(mockapi mockapi.MockApi, transactionRepository repository.TransactionRepository, walletRepository repository.WalletRepository) TransactionService {
	return &TransactionServiceImpl{
		MockApi:               mockapi,
		TransactionRepository: transactionRepository,
		WalletRepository:      walletRepository,
	}
}
