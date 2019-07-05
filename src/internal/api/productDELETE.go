package api

import (
	"github.com/ZergsLaw/korean/src/internal/api/restapi/operations"
	"net/http"
)

func (api *service) productDELETE(params operations.ProductDELETEParams, principal *int) operations.ProductDELETEResponder {
	ctx := params.HTTPRequest.Context()

	err := api.storage.ProductDelete(ctx, int(params.ID))
	if err != nil {
		api.log.Warn(err)

		return operations.NewProductDELETEDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewProductDELETEOK()
}
