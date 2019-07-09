// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewServiceStorageAPI creates a new ServiceStorage instance
func NewServiceStorageAPI(spec *loads.Document) *ServiceStorageAPI {
	return &ServiceStorageAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		BrandDELETEHandler: BrandDELETEHandlerFunc(func(params BrandDELETEParams, principal *int) BrandDELETEResponder {
			// return middleware.NotImplemented("operation BrandDELETE has not yet been implemented")
			return BrandDELETENotImplemented()
		}),
		BrandListHandler: BrandListHandlerFunc(func(params BrandListParams) BrandListResponder {
			// return middleware.NotImplemented("operation BrandList has not yet been implemented")
			return BrandListNotImplemented()
		}),
		BrandPOSTHandler: BrandPOSTHandlerFunc(func(params BrandPOSTParams, principal *int) BrandPOSTResponder {
			// return middleware.NotImplemented("operation BrandPOST has not yet been implemented")
			return BrandPOSTNotImplemented()
		}),
		GetUserHandler: GetUserHandlerFunc(func(params GetUserParams, principal *int) GetUserResponder {
			// return middleware.NotImplemented("operation GetUser has not yet been implemented")
			return GetUserNotImplemented()
		}),
		ListProductHandler: ListProductHandlerFunc(func(params ListProductParams) ListProductResponder {
			// return middleware.NotImplemented("operation ListProduct has not yet been implemented")
			return ListProductNotImplemented()
		}),
		LoginHandler: LoginHandlerFunc(func(params LoginParams) LoginResponder {
			// return middleware.NotImplemented("operation Login has not yet been implemented")
			return LoginNotImplemented()
		}),
		ProductHandler: ProductHandlerFunc(func(params ProductParams) ProductResponder {
			// return middleware.NotImplemented("operation Product has not yet been implemented")
			return ProductNotImplemented()
		}),
		ProductDELETEHandler: ProductDELETEHandlerFunc(func(params ProductDELETEParams, principal *int) ProductDELETEResponder {
			// return middleware.NotImplemented("operation ProductDELETE has not yet been implemented")
			return ProductDELETENotImplemented()
		}),
		ProductPOSTHandler: ProductPOSTHandlerFunc(func(params ProductPOSTParams, principal *int) ProductPOSTResponder {
			// return middleware.NotImplemented("operation ProductPOST has not yet been implemented")
			return ProductPOSTNotImplemented()
		}),
		UploadAvatarProductHandler: UploadAvatarProductHandlerFunc(func(params UploadAvatarProductParams, principal *int) UploadAvatarProductResponder {
			// return middleware.NotImplemented("operation UploadAvatarProduct has not yet been implemented")
			return UploadAvatarProductNotImplemented()
		}),

		// Applies when the "AdminCookie" header is set
		APIKeyAuth: func(token string) (*int, error) {
			return nil, errors.NotImplemented("api key auth (apiKey) AdminCookie from header param [AdminCookie] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ServiceStorageAPI Service storage brands and cosmetics. */
type ServiceStorageAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key AdminCookie provided in the header
	APIKeyAuth func(string) (*int, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// BrandDELETEHandler sets the operation handler for the brand d e l e t e operation
	BrandDELETEHandler BrandDELETEHandler
	// BrandListHandler sets the operation handler for the brand list operation
	BrandListHandler BrandListHandler
	// BrandPOSTHandler sets the operation handler for the brand p o s t operation
	BrandPOSTHandler BrandPOSTHandler
	// GetUserHandler sets the operation handler for the get user operation
	GetUserHandler GetUserHandler
	// ListProductHandler sets the operation handler for the list product operation
	ListProductHandler ListProductHandler
	// LoginHandler sets the operation handler for the login operation
	LoginHandler LoginHandler
	// ProductHandler sets the operation handler for the product operation
	ProductHandler ProductHandler
	// ProductDELETEHandler sets the operation handler for the product d e l e t e operation
	ProductDELETEHandler ProductDELETEHandler
	// ProductPOSTHandler sets the operation handler for the product p o s t operation
	ProductPOSTHandler ProductPOSTHandler
	// UploadAvatarProductHandler sets the operation handler for the upload avatar product operation
	UploadAvatarProductHandler UploadAvatarProductHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ServiceStorageAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ServiceStorageAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ServiceStorageAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ServiceStorageAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ServiceStorageAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ServiceStorageAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ServiceStorageAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ServiceStorageAPI
func (o *ServiceStorageAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "AdminCookieAuth")
	}

	if o.BrandDELETEHandler == nil {
		unregistered = append(unregistered, "BrandDELETEHandler")
	}

	if o.BrandListHandler == nil {
		unregistered = append(unregistered, "BrandListHandler")
	}

	if o.BrandPOSTHandler == nil {
		unregistered = append(unregistered, "BrandPOSTHandler")
	}

	if o.GetUserHandler == nil {
		unregistered = append(unregistered, "GetUserHandler")
	}

	if o.ListProductHandler == nil {
		unregistered = append(unregistered, "ListProductHandler")
	}

	if o.LoginHandler == nil {
		unregistered = append(unregistered, "LoginHandler")
	}

	if o.ProductHandler == nil {
		unregistered = append(unregistered, "ProductHandler")
	}

	if o.ProductDELETEHandler == nil {
		unregistered = append(unregistered, "ProductDELETEHandler")
	}

	if o.ProductPOSTHandler == nil {
		unregistered = append(unregistered, "ProductPOSTHandler")
	}

	if o.UploadAvatarProductHandler == nil {
		unregistered = append(unregistered, "UploadAvatarProductHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ServiceStorageAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ServiceStorageAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "apiKey":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.APIKeyAuth(token)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *ServiceStorageAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *ServiceStorageAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ServiceStorageAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ServiceStorageAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the service storage API
func (o *ServiceStorageAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ServiceStorageAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/brand"] = NewBrandDELETE(o.context, o.BrandDELETEHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/brand"] = NewBrandList(o.context, o.BrandListHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/brand"] = NewBrandPOST(o.context, o.BrandPOSTHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user"] = NewGetUser(o.context, o.GetUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/products"] = NewListProduct(o.context, o.ListProductHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = NewLogin(o.context, o.LoginHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/product/{productID}"] = NewProduct(o.context, o.ProductHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/product"] = NewProductDELETE(o.context, o.ProductDELETEHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/product"] = NewProductPOST(o.context, o.ProductPOSTHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/avatar/{productID}"] = NewUploadAvatarProduct(o.context, o.UploadAvatarProductHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ServiceStorageAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ServiceStorageAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ServiceStorageAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ServiceStorageAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
