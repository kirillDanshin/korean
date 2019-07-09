package api

import (
	"github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"net/http"
)

func (api *service) brandGET(params operations.BrandListParams) operations.BrandListResponder {
	ctx := params.HTTPRequest.Context()

	brands, err := api.storage.GetBrands(ctx)
	if err != nil {
		api.log.Warn(err)

		return operations.NewBrandListDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewBrandListOK().WithPayload(convertBrands(brands))
}
