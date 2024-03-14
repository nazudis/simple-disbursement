package dto

type GetAccessTokenRes struct {
	AccessToken string `json:"access_token"`
	IssuedAt    string `json:"issued_at"`
	ExpiresAt   string `json:"expires_at"`
}
