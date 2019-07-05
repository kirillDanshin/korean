package api

import (
	operations2 "github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"net/http"
)

func (api *service) brandDELETE(params operations2.BrandDELETEParams, principal *int) operations2.BrandDELETEResponder {
	ctx := params.HTTPRequest.Context()

	err := api.storage.BrandDelete(ctx, int(params.ID))
	if err != nil {
		api.log.Warn(err)

		return operations2.NewBrandDELETEDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations2.NewBrandDELETEOK()
}
