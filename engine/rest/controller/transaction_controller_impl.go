package controller

import (
	"net/http"

	"github.com/nazudis/disbursement/engine/rest/dto"
	"github.com/nazudis/disbursement/engine/rest/middleware"
	"github.com/nazudis/disbursement/engine/rest/transformer"
	"github.com/nazudis/disbursement/src/helper"
	"github.com/nazudis/disbursement/src/service"
	"github.com/nazudis/netper"
)

type TransactionControllerImpl struct {
	Service     service.TransactionService
	Transformer transformer.TransactionTransformer
}

// GetAccountVerification implements TransactionController.
func (c *TransactionControllerImpl) GetAccountVerification(w http.ResponseWriter, r *http.Request) {
	res := netper.PlugResponse(w)
	req := netper.PlugRequest(r, w)

	var (
		accountNumber = req.GetString("account_number")
		bankShortCode = req.GetString("bank_short_code")
	)

	if accountNumber == "" || bankShortCode == "" {
		resData := helper.ReplyFail(http.StatusBadRequest, "account number and bank short code are required")
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	acc, err := c.Service.GetAccountVerification(accountNumber, bankShortCode)
	if err != nil {
		resData := helper.ReplyFail(http.StatusBadRequest, err.Error())
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	data := c.Transformer.TransformGetAccountVerification(acc)
	resData := helper.ReplySuccess("success verify account", data)
	_ = res.ReplyCustom(resData.Status, resData)
}

// PostDisbursement implements TransactionController.
func (c *TransactionControllerImpl) PostDisbursement(w http.ResponseWriter, r *http.Request) {
	res := netper.PlugResponse(w)
	req := netper.PlugRequest(r, w)

	params, err := netper.ParseTo[dto.PostDisbursementReq](req)
	if err != nil {
		resData := helper.ReplyFail(http.StatusBadRequest, err.Error())
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	authSession := middleware.GetAuthSessionFromCtx(r.Context())

	trx, err := c.Service.PostDisbursement(service.PostDisbursementParams{
		OwnerID:               authSession.OwnerID.String(),
		ReferenceID:           params.ReferenceID,
		Description:           params.Description,
		Amount:                params.Amount,
		Type:                  params.Type,
		BankShortCode:         params.BankShortCode,
		BankAccountNo:         params.BankAccountNo,
		BankAccountHolderName: params.BankAccountHolderName,
	})
	if err != nil {
		resData := helper.ReplyFail(http.StatusBadRequest, err.Error())
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	data := c.Transformer.TransformDisbursement(trx)
	resData := helper.ReplySuccess("your request is being process", data)
	_ = res.ReplyCustom(resData.Status, resData)
}

// PostCallbackDisbursement implements TransactionController.
func (c *TransactionControllerImpl) PostCallbackDisbursement(w http.ResponseWriter, r *http.Request) {
	res := netper.PlugResponse(w)
	req := netper.PlugRequest(r, w)

	params, err := netper.ParseTo[dto.PostCallbackDisbursementReq](req)
	if err != nil {
		resData := helper.ReplyFail(http.StatusBadRequest, err.Error())
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	err = c.Service.PostCallbackDisbursement(service.PostCallbackDisbursementParams{
		ID:                    params.ID,
		Status:                params.Status,
		CreatedAt:             params.CreatedAt,
		ReferenceID:           params.ReferenceID,
		Description:           params.Description,
		Amount:                params.Amount,
		Type:                  params.Type,
		BankName:              params.BankName,
		BankShortCode:         params.BankShortCode,
		BankAccountNo:         params.BankAccountNo,
		BankAccountHolderName: params.BankAccountHolderName,
	})
	if err != nil {
		resData := helper.ReplyFail(http.StatusBadRequest, err.Error())
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	resData := helper.ReplySuccess("success", nil)
	_ = res.ReplyCustom(resData.Status, resData)
}

func NewTransactionController(service service.TransactionService, transformer transformer.TransactionTransformer) TransactionController {
	return &TransactionControllerImpl{
		Service:     service,
		Transformer: transformer,
	}
}
