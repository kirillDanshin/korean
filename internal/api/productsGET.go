package api

import (
	"github.com/ZergsLaw/korean/internal/api/restapi/operations"
	"github.com/ZergsLaw/korean/internal/db"
	"net/http"
)

func (api *service) productsGET(params operations.ListProductParams) operations.ListProductResponder {
	ctx := params.HTTPRequest.Context()

	queryParams := db.Params{
		BrandID: intPoint(params.Brand),
		Query:   stringPoint(params.Query),
		Pagination: db.Pagination{
			Limit:  intPoint(params.Limit),
			Offset: intPoint(params.Offset),
		},
	}

	products, err := api.storage.ListProduct(ctx, queryParams)
	if err != nil {
		api.log.Warn(err)

		return operations.NewListProductDefault(http.StatusInternalServerError).
			WithPayload(createErr(http.StatusInternalServerError))
	}

	return operations.NewListProductOK().WithPayload(convertArrayProduct(products))
}

func intPoint(point *int64) int {
	if point != nil {
		return int(*point)
	}
	return 0
}

func stringPoint(point *string) string {
	if point != nil {
		return *point
	}
	return ""
}
