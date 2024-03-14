package transformer

import (
	"github.com/nazudis/disbursement/engine/rest/dto"
	"github.com/nazudis/disbursement/src/model/entity"
)

type AuthTransformer interface {
	TransformGetAccessToken(authSession *entity.AuthSession) dto.GetAccessTokenRes
}
