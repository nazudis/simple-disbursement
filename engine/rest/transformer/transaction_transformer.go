package transformer

import (
	"github.com/nazudis/disbursement/engine/rest/dto"
	mockapiResponse "github.com/nazudis/disbursement/src/dependencies/mockapi/response"
	"github.com/nazudis/disbursement/src/model/entity"
)

type TransactionTransformer interface {
	TransformGetAccountVerification(*mockapiResponse.Account) dto.GetAccountVerificationRes
	TransformDisbursement(*entity.Transaction) dto.PostDisbursementRes
}
