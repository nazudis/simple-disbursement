package request

type PostTransactionRequest struct {
	ReferenceID   string  `json:"reference_id,omitempty"`
	Type          string  `json:"type,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	AccountNumber string  `json:"account_number,omitempty"`
}
