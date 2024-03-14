package transformer

import (
	"github.com/nazudis/disbursement/engine/rest/dto"
	mockapiResponse "github.com/nazudis/disbursement/src/dependencies/mockapi/response"
	"github.com/nazudis/disbursement/src/model/entity"
)

type TransactionTransformerImpl struct{}

// TransformGetAccountVerification implements TransactionTransformer.
func (t *TransactionTransformerImpl) TransformGetAccountVerification(acc *mockapiResponse.Account) dto.GetAccountVerificationRes {
	data := dto.GetAccountVerificationRes{
		AccountNumber: acc.AccountNumber,
		AccountName:   acc.AccountName,
		BankShortCode: acc.BankShortCode,
	}
	return data
}

// TransformDisbursement implements TransactionTransformer.
func (t *TransactionTransformerImpl) TransformDisbursement(trx *entity.Transaction) dto.PostDisbursementRes {
	data := dto.PostDisbursementRes{
		ID:                     trx.ID.String(),
		Type:                   trx.Type,
		Amount:                 trx.Amount,
		ReferenceID:            trx.ReferenceID,
		RecipientAccountNumber: trx.RecipientAccountNumber,
		Status:                 trx.Status,
		Description:            trx.Description,
	}
	return data
}

func NewTransactionTransformer() TransactionTransformer {
	return &TransactionTransformerImpl{}
}
