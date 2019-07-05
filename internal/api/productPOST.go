package api

import (
	operations2 "github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"github.com/ZergsLaw/korean/internal/db"
	"net/http"
)

func (api *service) productPOST(params operations2.ProductPOSTParams, principal *int) operations2.ProductPOSTResponder {
	ctx := params.HTTPRequest.Context()

	newProduct := db.NewProduct{
		Name:        *params.Args.Name,
		Description: *params.Args.Description,
		Apply:       *params.Args.Apply,
		Price:       int(*params.Args.Price),
		BrandID:     int(*params.Args.BrandID),
	}

	product, err := api.storage.ProductCreate(ctx, newProduct)
	if err != nil {
		api.log.Warn(err)

		return operations2.NewProductPOSTDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations2.NewProductPOSTOK().WithPayload(convertProduct(product))
}
