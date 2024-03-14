package transformer

import (
	"time"

	"github.com/nazudis/disbursement/engine/rest/dto"
	"github.com/nazudis/disbursement/src/model/entity"
)

type AuthTransformerImpl struct{}

// TransformGetAccessToken implements AuthTransformer.
func (t *AuthTransformerImpl) TransformGetAccessToken(authSession *entity.AuthSession) dto.GetAccessTokenRes {
	data := dto.GetAccessTokenRes{
		AccessToken: authSession.Token,
		IssuedAt:    authSession.IssuedAt.Format(time.RFC3339),
		ExpiresAt:   authSession.ExpiredAt.Format(time.RFC3339),
	}
	return data
}

func NewAuthTransformer() AuthTransformer {
	return &AuthTransformerImpl{}
}
