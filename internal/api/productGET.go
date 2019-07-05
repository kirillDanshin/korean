package api

import (
	operations2 "github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"net/http"
)

func (api *service) productGET(params operations2.ProductParams) operations2.ProductResponder {
	ctx := params.HTTPRequest.Context()

	product, err := api.storage.ProductByID(ctx, int(params.ProductID))
	if err != nil {
		api.log.Warn(err)

		return operations2.NewProductDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations2.NewProductOK().WithPayload(convertProduct(product))
}
