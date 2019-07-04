package api

import (
	"net/http"
	"storage/internal/api/models"
	"storage/internal/api/restapi/operations"
	"storage/internal/db"
)

func (api *service) brandPOST(params operations.BrandPOSTParams, principal *int) operations.BrandPOSTResponder {
	ctx := params.HTTPRequest.Context()

	ID, err := api.storage.BrandCreate(ctx, db.NewBrand{Name: *params.Args.Name})
	if err != nil {
		api.log.Warn(err)

		return operations.NewBrandPOSTDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewBrandPOSTOK().WithPayload(&models.Success{ID: models.ID(ID)})
}
