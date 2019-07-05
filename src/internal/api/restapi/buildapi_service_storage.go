// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"io"
	"net/http"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"

	"github.com/ZergsLaw/korean/src/internal/api/restapi/operations"
)

func BuildAPI(
	swaggerSpec *loads.Document,
	ServeError func(http.ResponseWriter, *http.Request, error),
	Logger func(string, ...interface{}),

	JSONConsumer func(r io.Reader, target interface{}) error,

	JSONProducer func(w io.Writer, data interface{}) error,

	APIKeyAuth func(token string) (*int, error),
	APIAuthorizer runtime.Authorizer,

	BrandDELETE func(params operations.BrandDELETEParams, principal *int) operations.BrandDELETEResponder,
	BrandList func(params operations.BrandListParams, principal *int) operations.BrandListResponder,
	BrandPOST func(params operations.BrandPOSTParams, principal *int) operations.BrandPOSTResponder,
	ListProduct func(params operations.ListProductParams) operations.ListProductResponder,
	Login func(params operations.LoginParams) operations.LoginResponder,
	Product func(params operations.ProductParams) operations.ProductResponder,
	ProductDELETE func(params operations.ProductDELETEParams, principal *int) operations.ProductDELETEResponder,
	ProductPOST func(params operations.ProductPOSTParams, principal *int) operations.ProductPOSTResponder,

	ServerShutdown func(),

) *operations.ServiceStorageAPI {
	api := operations.NewServiceStorageAPI(swaggerSpec)

	if ServeError != nil {
		api.ServeError = errors.ServeError
	}

	if Logger != nil {
		api.Logger = Logger
	}

	if JSONConsumer != nil {
		api.JSONConsumer = runtime.ConsumerFunc(JSONConsumer)
	}

	if JSONProducer != nil {
		api.JSONProducer = runtime.ProducerFunc(JSONProducer)
	}

	if APIKeyAuth != nil {
		api.APIKeyAuth = APIKeyAuth
	}
	if APIAuthorizer != nil {
		api.APIAuthorizer = APIAuthorizer
	}

	if BrandDELETE != nil {
		api.BrandDELETEHandler = operations.BrandDELETEHandlerFunc(BrandDELETE)
	}

	if BrandList != nil {
		api.BrandListHandler = operations.BrandListHandlerFunc(BrandList)
	}

	if BrandPOST != nil {
		api.BrandPOSTHandler = operations.BrandPOSTHandlerFunc(BrandPOST)
	}

	if ListProduct != nil {
		api.ListProductHandler = operations.ListProductHandlerFunc(ListProduct)
	}

	if Login != nil {
		api.LoginHandler = operations.LoginHandlerFunc(Login)
	}

	if Product != nil {
		api.ProductHandler = operations.ProductHandlerFunc(Product)
	}

	if ProductDELETE != nil {
		api.ProductDELETEHandler = operations.ProductDELETEHandlerFunc(ProductDELETE)
	}

	if ProductPOST != nil {
		api.ProductPOSTHandler = operations.ProductPOSTHandlerFunc(ProductPOST)
	}

	if ServerShutdown != nil {
		api.ServerShutdown = ServerShutdown
	}

	return api
}
