package mockapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nazudis/disbursement/src/dependencies/mockapi/request"
	"github.com/nazudis/disbursement/src/dependencies/mockapi/response"
)

type MockApiImpl struct {
	host string
}

// PostTransaction implements MockApi.
func (p *MockApiImpl) PostTransaction(data request.PostTransactionRequest) (*response.PostTransactionResponse, error) {
	path := "/transactions"

	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := p.request(path, http.MethodPost, jsonBody, nil)
	if err != nil {
		return nil, err
	}

	defer p.closeBody(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to post transaction: %s", string(respBody))
	}

	trx, err := readBody[response.PostTransactionResponse](respBody)
	if err != nil {

		return nil, err
	}

	return trx, nil
}

// GetAccountByAccountNumberAndBankShortCode implements PlatformAPI.
func (p *MockApiImpl) GetAccountByAccountNumberAndBankShortCode(accountNumber string, bankShortCode string) (*response.Account, error) {
	path := "/accounts"
	params := []Param{
		{Key: "account_number", Value: accountNumber},
		{Key: "bank_short_code", Value: bankShortCode},
	}

	resp, err := p.request(path, http.MethodGet, nil, params)
	if err != nil {
		return nil, err
	}

	defer p.closeBody(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get account: %s", string(respBody))
	}

	accounts, err := readBody[[]response.Account](respBody)
	if err != nil {
		return nil, err
	}

	if len(*accounts) == 0 {
		return nil, fmt.Errorf("account not found")
	}

	return &(*accounts)[0], nil
}

func readBody[T any](respBody []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// closeBody implements PlatformAPI.
func (*MockApiImpl) closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Panic(err)
	}
}

// request implements PlatformAPI
func (p *MockApiImpl) request(path string, method string, body []byte, params []Param) (res *http.Response, err error) {
	client := http.Client{}

	url := fmt.Sprintf("%s%s", p.host, path)
	data := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	// Set Query Param URL
	httpParam := req.URL.Query()
	for _, param := range params {
		httpParam.Add(param.Key, param.Value)
	}

	// Set Header
	req.Header = http.Header{}
	if method == http.MethodPost {
		req.Header.Add("content-type", "application/json")
	}
	req.URL.RawQuery = httpParam.Encode()

	res, err = client.Do(req)
	return
}

func NewMockApi() MockApi {
	return &MockApiImpl{
		host: "https://65f1c094034bdbecc76396ac.mockapi.io/api/v1",
	}
}
