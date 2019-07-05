package api

import (
	models2 "github.com/ZergsLaw/korean/internal/api/models"
	operations2 "github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"github.com/ZergsLaw/korean/internal/db"
	"github.com/go-openapi/swag"
	"net/http"
)

func (api *service) brandPOST(params operations2.BrandPOSTParams, principal *int) operations2.BrandPOSTResponder {
	ctx := params.HTTPRequest.Context()

	brands, err := api.storage.BrandCreate(ctx, db.NewBrand{Name: *params.Args.Name})
	if err != nil {
		api.log.Warn(err)

		return operations2.NewBrandPOSTDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations2.NewBrandPOSTOK().WithPayload(convertBrands(brands))
}

func convertBrands(brands []db.Brand) []*models2.Brand {
	convertsValue := make([]*models2.Brand, len(brands))
	for i := range brands {
		convertsValue[i] = &models2.Brand{ID: models2.ID(brands[i].Id), Name: swag.String(brands[i].Name)}
	}
	return convertsValue
}
