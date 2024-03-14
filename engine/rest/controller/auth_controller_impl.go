package controller

import (
	"net/http"

	"github.com/nazudis/disbursement/engine/rest/transformer"
	"github.com/nazudis/disbursement/src/helper"
	"github.com/nazudis/disbursement/src/service"
	"github.com/nazudis/netper"
)

type AuthControllerImpl struct {
	Service     service.AuthService
	Transformer transformer.AuthTransformer
}

// GetAccessToken implements AuthController.
func (c *AuthControllerImpl) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	res := netper.PlugResponse(w)
	req := netper.PlugRequest(r, w)

	if !req.HasUser() {
		resData := helper.ReplyFail(http.StatusUnauthorized, "unauthorized. credential required")
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	if req.GetUsername() == "" || req.GetPassword() == "" {
		resData := helper.ReplyFail(http.StatusBadRequest, "username (client_id) and password (client_secret) is required")
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	authSession, err := c.Service.GetAccessToken(req.GetUsername(), req.GetPassword())
	if err != nil {
		resData := helper.ReplyFail(http.StatusBadRequest, err.Error())
		_ = res.ReplyCustom(resData.Status, resData)
		return
	}

	data := c.Transformer.TransformGetAccessToken(authSession)
	resData := helper.ReplySuccess("success get access token", data)
	_ = res.ReplyCustom(resData.Status, resData)
}

func NewAuthController(service service.AuthService, transformer transformer.AuthTransformer) AuthController {
	return &AuthControllerImpl{Service: service, Transformer: transformer}
}
