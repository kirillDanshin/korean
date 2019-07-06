package api

import "github.com/ZergsLaw/korean/internal/api/restapi/operations"

func (api *service) avatarPUT(params operations.UploadAvatarProductParams, principal *int) operations.UploadAvatarProductResponder  {
	return operations.NewUploadAvatarProductOK()
}