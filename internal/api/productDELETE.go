package api

import (
	operations2 "github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"net/http"
)

func (api *service) productDELETE(params operations2.ProductDELETEParams, principal *int) operations2.ProductDELETEResponder {
	ctx := params.HTTPRequest.Context()

	err := api.storage.ProductDelete(ctx, int(params.ID))
	if err != nil {
		api.log.Warn(err)

		return operations2.NewProductDELETEDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations2.NewProductDELETEOK()
}
