package api

import (
	"github.com/go-openapi/swag"
	"net/http"
	"storage/internal/api/models"
	"storage/internal/api/restapi/operations"
	"storage/internal/lib/hash"
)

func (api *service) login(params operations.LoginParams) operations.LoginResponder {
	ctx := params.HTTPRequest.Context()

	token := hash.RandToken()
	if api.adminUsername != string(params.Args.Username) || api.adminPass != string(params.Args.Password) {
		return operations.NewLoginDefault(http.StatusUnauthorized).WithPayload(createErr(http.StatusUnauthorized))
	}

	err := api.session.Set(ctx, token, 1)
	if err != nil {
		api.log.Warn(err)

		return operations.NewLoginDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewLoginOK().WithAdminCookie(token).WithPayload(&models.UserInfo{
		Token:    swag.String(token),
		Username: models.Username(api.adminUsername),
	})
}
