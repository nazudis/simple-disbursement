package middleware

import (
	"context"
	"net/http"

	"github.com/nazudis/disbursement/src/database"
	"github.com/nazudis/disbursement/src/helper"
	"github.com/nazudis/disbursement/src/model/entity"
	"github.com/nazudis/disbursement/src/repository"
	"github.com/nazudis/netper"
)

type CtxKey string

const AuthCtxKey CtxKey = "auth_session"

func AuthSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authSessionRepository := repository.NewAuthSessionRepository(database.GetDB())

		res := netper.PlugResponse(w)
		req := netper.TouchRequest(r, w)

		if !req.HasHeader("Authorization") {
			resData := helper.ReplyFail(http.StatusForbidden, "Forbidden access! Missing token in header")
			_ = res.ReplyCustom(resData.Status, resData)
			return
		}

		prefix := "Bearer "
		auth := req.Header("Authorization")

		if len(auth) < len(prefix) || auth[:len(prefix)] != prefix {
			resData := helper.ReplyFail(http.StatusUnauthorized, "Unauthorized! Invalid token format")
			_ = res.ReplyCustom(resData.Status, resData)
			return
		}

		token := auth[len(prefix):]

		jwtClaims, err := helper.ParseJWTToken(token)
		if err != nil {
			resData := helper.ReplyFail(http.StatusBadRequest, err.Error())
			_ = res.ReplyCustom(resData.Status, resData)
			return
		}

		authSession, err := authSessionRepository.FirstByID(jwtClaims["auth_id"].(string))
		if err != nil {
			resData := helper.ReplyFail(http.StatusUnprocessableEntity, err.Error())
			_ = res.ReplyCustom(resData.Status, resData)
			return
		}
		if authSession == nil {
			resData := helper.ReplyFail(http.StatusUnauthorized, "Unauthorized! Invalid token")
			_ = res.ReplyCustom(resData.Status, resData)
			return
		}
		if authSession.IsExpired || authSession.IsRevoked {
			resData := helper.ReplyFail(http.StatusUnauthorized, "Unauthorized! Token has been expired or revoked")
			_ = res.ReplyCustom(resData.Status, resData)
			return
		}

		// Set the auth session is revoked after use
		err = authSessionRepository.Revoked(authSession)
		if err != nil {
			resData := helper.ReplyFail(http.StatusUnprocessableEntity, err.Error())
			_ = res.ReplyCustom(resData.Status, resData)
			return
		}

		// Set the auth session to the request context
		ctx := context.WithValue(r.Context(), AuthCtxKey, authSession)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetAuthSessionFromCtx(ctx context.Context) *entity.AuthSession {
	return ctx.Value(AuthCtxKey).(*entity.AuthSession)
}
