package api

import (
	"github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"net/http"
)

func (api *service) avatarPUT(params operations.UploadAvatarProductParams, principal *int) operations.UploadAvatarProductResponder {
	ctx := params.HTTPRequest.Context()

	err := api.storage.SetAvatar(ctx, int(params.ProductID), params.File)
	if err != nil {
		api.log.Warn(err)

		return operations.NewUploadAvatarProductDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewUploadAvatarProductOK()
}
