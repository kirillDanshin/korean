// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"storage/internal/api/restapi/operations"
)

//go:generate swagger generate server --target ../../api --name ServiceStorage --spec ../swagger.yml --principal int --exclude-main --strict

func configureFlags(api *operations.ServiceStorageAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ServiceStorageAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "AdminCookie" header is set
	api.APIKeyAuth = func(token string) (*int, error) {
		return nil, errors.NotImplemented("api key auth (apiKey) AdminCookie from header param [AdminCookie] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.BrandDELETEHandler = operations.BrandDELETEHandlerFunc(func(params operations.BrandDELETEParams, principal *int) operations.BrandDELETEResponder {
		return operations.BrandDELETENotImplemented()
	})
	api.BrandPOSTHandler = operations.BrandPOSTHandlerFunc(func(params operations.BrandPOSTParams, principal *int) operations.BrandPOSTResponder {
		return operations.BrandPOSTNotImplemented()
	})
	api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) operations.LoginResponder {
		return operations.LoginNotImplemented()
	})
	api.ProductDELETEHandler = operations.ProductDELETEHandlerFunc(func(params operations.ProductDELETEParams, principal *int) operations.ProductDELETEResponder {
		return operations.ProductDELETENotImplemented()
	})
	api.ProductPOSTHandler = operations.ProductPOSTHandlerFunc(func(params operations.ProductPOSTParams, principal *int) operations.ProductPOSTResponder {
		return operations.ProductPOSTNotImplemented()
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
