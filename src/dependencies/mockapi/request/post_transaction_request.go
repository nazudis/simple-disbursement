package request

type PostTransactionRequest struct {
	Type          string  `json:"type,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	AccountNumber string  `json:"account_number,omitempty"`
}
