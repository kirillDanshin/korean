package api

import (
	 "github.com/ZergsLaw/korean/internal/api/models"
	 "github.com/ZergsLaw/korean/internal/api/restapi/operations"
	 "github.com/ZergsLaw/korean/internal/lib/hash"
	"net/http"

	"github.com/go-openapi/swag"
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
