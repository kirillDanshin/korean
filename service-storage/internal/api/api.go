package api

import (
	"net"
	"net/http"
	"storage/internal/api/models"
	"storage/internal/api/restapi"
	"storage/internal/db"
	"storage/internal/session"
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

		storage Storage
		session session.Store
		log     Log
	}
	// Storage - main store products and brands.
	Storage interface {
		BrandCreate(ctx Ctx, brand db.NewBrand) (ID int, err error)
		BrandDelete(ctx Ctx, brandID int) (err error)

		ProductCreate(ctx Ctx, product db.NewProduct) (ID int, err error)
		ProductDelete(ctx Ctx, productID int) (err error)
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
func Serve(log Log, store Storage, sessionStore session.Store, cfg Configuration) (*restapi.Server, error) {
	svc := service{
		adminPass:     cfg.AdminPass,
		adminUsername: cfg.AdminUsername,
		storage:       store,
		log:           log,
		session:       sessionStore,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrapf(err, "failed loads embedded")
	}

	searchAPI := restapi.BuildAPI(
		swaggerSpec,
		nil,
		log.Printf,
		nil,
		nil,
		sessionStore.GetID,
		security.Authorized(),
		svc.brandDELETE,
		svc.brandPOST,
		svc.login,
		svc.productDELETE,
		svc.productPOST,
		nil,
	)

	globalMiddleware := func(handler http.Handler) http.Handler {
		logger := makeLogger(swaggerSpec.BasePath())
		return logger(recovery(handleCORS(handler)))
	}

	server := restapi.NewServer(searchAPI)
	server.SetHandler(globalMiddleware(searchAPI.Serve(accessLog)))
	server.Port = cfg.Port
	server.Host = cfg.Host

	log.Info("protocol", "version", swaggerSpec.Spec().Info.Version, "address", net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)))
	return server, nil
}

func createErr(code int) *models.Error {
	return &models.Error{
		Code:    swag.Int32(int32(code)),
		Message: swag.String(http.StatusText(code)),
	}
}
