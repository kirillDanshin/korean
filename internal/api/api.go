package api

import (
	models2 "github.com/ZergsLaw/korean/internal/api/models"
	restapi2 "github.com/ZergsLaw/korean/internal/api/restapi"
	"github.com/ZergsLaw/korean/internal/db"
	"github.com/ZergsLaw/korean/internal/session"
	"net"
	"net/http"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"golang.org/x/net/context"
)

type (
	service struct {
		adminUsername string
		adminPass     string

		storage db.Storage
		session session.Store
		log     Log
	}

	// Configuration contains config for api service.
	Configuration struct {
		AdminUsername string
		AdminPass     string

		Port int
		Host string
	}

	// Ctx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

// Serve must be called once before using this package.
func Serve(log Log, store db.Storage, sessionStore session.Store, cfg Configuration) (*restapi2.Server, error) {
	svc := service{
		adminPass:     cfg.AdminPass,
		adminUsername: cfg.AdminUsername,
		storage:       store,
		log:           log,
		session:       sessionStore,
	}

	swaggerSpec, err := loads.Embedded(restapi2.SwaggerJSON, restapi2.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrapf(err, "failed loads embedded")
	}

	searchAPI := restapi2.BuildAPI(
		swaggerSpec,
		nil,
		log.Printf,
		nil,
		nil,
		sessionStore.GetID,
		security.Authorized(),
		svc.brandDELETE,
		svc.brandGET,
		svc.brandPOST,
		svc.productsGET,
		svc.login,
		svc.productGET,
		svc.productDELETE,
		svc.productPOST,
		nil,
	)

	globalMiddleware := func(handler http.Handler) http.Handler {
		logger := makeLogger(swaggerSpec.BasePath())
		return logger(recovery(handleCORS(handler)))
	}

	server := restapi2.NewServer(searchAPI)
	server.SetHandler(globalMiddleware(searchAPI.Serve(accessLog)))
	server.Port = cfg.Port
	server.Host = cfg.Host

	log.Info("protocol", "version", swaggerSpec.Spec().Info.Version, "address", net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)))
	return server, nil
}

func createErr(code int) *models2.Error {
	return &models2.Error{
		Code:    swag.Int32(int32(code)),
		Message: swag.String(http.StatusText(code)),
	}
}

func convertArrayProduct(products []db.Product) []*models2.Product {
	convertProducts := make([]*models2.Product, len(products))

	for i := range products {
		convertProducts[i] = convertProduct(&products[i])
	}

	return convertProducts
}

func convertProduct(product *db.Product) *models2.Product {
	return &models2.Product{
		Apply:       swag.String(product.Apply),
		BrandName:   swag.String(product.Brand),
		Description: swag.String(product.Description),
		ID:          models2.ID(product.ID),
		Name:        swag.String(product.Name),
		Price:       swag.Int64(int64(product.Price)),
		AvatarURL:   product.Avatar.String,
	}
}
