package api

import (
	operations2 "github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"net/http"
)

func (api *service) brandGET(params operations2.BrandListParams, principal *int) operations2.BrandListResponder {
	ctx := params.HTTPRequest.Context()

	brands, err := api.storage.GetBrands(ctx)
	if err != nil {
		api.log.Warn(err)

		return operations2.NewBrandListDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations2.NewBrandListOK().WithPayload(convertBrands(brands))
}
