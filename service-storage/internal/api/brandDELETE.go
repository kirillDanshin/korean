package api

import (
	"net/http"
	"storage/internal/api/restapi/operations"
)

func (api *service) brandDELETE(params operations.BrandDELETEParams, principal *int) operations.BrandDELETEResponder {
	ctx := params.HTTPRequest.Context()

	err := api.storage.BrandDelete(ctx, int(params.ID))
	if err != nil {
		api.log.Warn(err)

		return operations.NewBrandDELETEDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewBrandDELETEOK()
}
