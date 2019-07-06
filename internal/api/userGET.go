package api

import (
	"github.com/ZergsLaw/korean/internal/api/models"
	"github.com/ZergsLaw/korean/internal/api/restapi/operations"
)

func (api *service) userGET(params operations.GetUserParams, principal *int) operations.GetUserResponder {
	return operations.NewGetUserOK().WithPayload(&models.UserInfo{Username: models.Username(api.adminUsername)})
}
