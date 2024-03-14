package response

type PostTransactionResponse struct {
	ID            string `json:"id"`
	TrxID         string `json:"trx_id"`
	ReferenceID   string `json:"reference_id"`
	Type          string `json:"type"`
	Amount        int    `json:"amount"`
	AccountNumber string `json:"account_number"`
	CreatedAt     string `json:"createdAt"`
}
