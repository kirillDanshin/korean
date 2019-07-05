package api

import (
	"github.com/ZergsLaw/korean/src/internal/api/restapi/operations"
	"net/http"
)

func (api *service) productGET(params operations.ProductParams) operations.ProductResponder {
	ctx := params.HTTPRequest.Context()

	product, err := api.storage.ProductByID(ctx, int(params.ProductID))
	if err != nil {
		api.log.Warn(err)

		return operations.NewProductDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewProductOK().WithPayload(convertProduct(product))
}
