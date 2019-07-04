// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/disheshub/culinario/internal/api/generated/restapi/operations"
	"github.com/disheshub/culinario/internal/app"
)

//go:generate swagger generate server --target ../../generated --name Culinario --spec ../swagger.yml --principal app.Session --exclude-main --strict

func configureFlags(api *operations.CulinarioAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CulinarioAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-culinario-token" header is set
	api.APIKeyAuth = func(token string) (*app.Session, error) {
		return nil, errors.NotImplemented("api key auth (apiKey) X-culinario-token from header param [X-culinario-token] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.CreateIngredientHandler = operations.CreateIngredientHandlerFunc(func(params operations.CreateIngredientParams, principal *app.Session) operations.CreateIngredientResponder {
		return operations.CreateIngredientNotImplemented()
	})
	api.CreateRecipeHandler = operations.CreateRecipeHandlerFunc(func(params operations.CreateRecipeParams, principal *app.Session) operations.CreateRecipeResponder {
		return operations.CreateRecipeNotImplemented()
	})
	api.ExistsEmailHandler = operations.ExistsEmailHandlerFunc(func(params operations.ExistsEmailParams) operations.ExistsEmailResponder {
		return operations.ExistsEmailNotImplemented()
	})
	api.GetIngredientsHandler = operations.GetIngredientsHandlerFunc(func(params operations.GetIngredientsParams, principal *app.Session) operations.GetIngredientsResponder {
		return operations.GetIngredientsNotImplemented()
	})
	api.GetRecipeHandler = operations.GetRecipeHandlerFunc(func(params operations.GetRecipeParams) operations.GetRecipeResponder {
		return operations.GetRecipeNotImplemented()
	})
	api.GetUserHandler = operations.GetUserHandlerFunc(func(params operations.GetUserParams) operations.GetUserResponder {
		return operations.GetUserNotImplemented()
	})
	api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) operations.LoginResponder {
		return operations.LoginNotImplemented()
	})
	api.NewMeasureWeightHandler = operations.NewMeasureWeightHandlerFunc(func(params operations.NewMeasureWeightParams, principal *app.Session) operations.NewMeasureWeightResponder {
		return operations.NewMeasureWeightNotImplemented()
	})
	api.RegistrationHandler = operations.RegistrationHandlerFunc(func(params operations.RegistrationParams) operations.RegistrationResponder {
		return operations.RegistrationNotImplemented()
	})
	api.RemoveIngredientHandler = operations.RemoveIngredientHandlerFunc(func(params operations.RemoveIngredientParams, principal *app.Session) operations.RemoveIngredientResponder {
		return operations.RemoveIngredientNotImplemented()
	})
	api.RemoveRecipeHandler = operations.RemoveRecipeHandlerFunc(func(params operations.RemoveRecipeParams, principal *app.Session) operations.RemoveRecipeResponder {
		return operations.RemoveRecipeNotImplemented()
	})
	api.RmMeasureWeightHandler = operations.RmMeasureWeightHandlerFunc(func(params operations.RmMeasureWeightParams, principal *app.Session) operations.RmMeasureWeightResponder {
		return operations.RmMeasureWeightNotImplemented()
	})
	api.UpdateUserInfoHandler = operations.UpdateUserInfoHandlerFunc(func(params operations.UpdateUserInfoParams, principal *app.Session) operations.UpdateUserInfoResponder {
		return operations.UpdateUserInfoNotImplemented()
	})
	api.UploadAvatarHandler = operations.UploadAvatarHandlerFunc(func(params operations.UploadAvatarParams, principal *app.Session) operations.UploadAvatarResponder {
		return operations.UploadAvatarNotImplemented()
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
