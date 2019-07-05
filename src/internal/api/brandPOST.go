package api

import (
	"github.com/ZergsLaw/korean/src/internal/api/models"
	"github.com/ZergsLaw/korean/src/internal/api/restapi/operations"
	"github.com/ZergsLaw/korean/src/internal/db"
	"github.com/go-openapi/swag"
	"net/http"
)

func (api *service) brandPOST(params operations.BrandPOSTParams, principal *int) operations.BrandPOSTResponder {
	ctx := params.HTTPRequest.Context()

	brands, err := api.storage.BrandCreate(ctx, db.NewBrand{Name: *params.Args.Name})
	if err != nil {
		api.log.Warn(err)

		return operations.NewBrandPOSTDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewBrandPOSTOK().WithPayload(convertBrands(brands))
}

func convertBrands(brands []db.Brand) []*models.Brand {
	convertsValue := make([]*models.Brand, len(brands))
	for i := range brands {
		convertsValue[i] = &models.Brand{ID: models.ID(brands[i].Id), Name: swag.String(brands[i].Name)}
	}
	return convertsValue
}
