package mockapi

import (
	"io"
	"net/http"

	"github.com/nazudis/disbursement/src/dependencies/mockapi/request"
	"github.com/nazudis/disbursement/src/dependencies/mockapi/response"
)

type MockApi interface {
	request(path string, method string, body []byte, params []Param) (response *http.Response, err error)
	closeBody(body io.ReadCloser)
	GetAccountByAccountNumberAndBankShortCode(accountNumber string, bankShortCode string) (*response.Account, error)
	PostTransaction(data request.PostTransactionRequest) (*response.PostTransactionResponse, error)
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
