package api

import (
	models2 "github.com/ZergsLaw/korean/internal/api/models"
	operations2 "github.com/ZergsLaw/korean/internal/api/restapi/operations"
	hash2 "github.com/ZergsLaw/korean/internal/lib/hash"
	"net/http"

	"github.com/go-openapi/swag"
)

func (api *service) login(params operations2.LoginParams) operations2.LoginResponder {
	ctx := params.HTTPRequest.Context()

	token := hash2.RandToken()
	if api.adminUsername != string(params.Args.Username) || api.adminPass != string(params.Args.Password) {
		return operations2.NewLoginDefault(http.StatusUnauthorized).WithPayload(createErr(http.StatusUnauthorized))
	}

	err := api.session.Set(ctx, token, 1)
	if err != nil {
		api.log.Warn(err)

		return operations2.NewLoginDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations2.NewLoginOK().WithAdminCookie(token).WithPayload(&models2.UserInfo{
		Token:    swag.String(token),
		Username: models2.Username(api.adminUsername),
	})
}
