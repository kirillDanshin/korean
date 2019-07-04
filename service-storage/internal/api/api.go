package api

import (
	"net"
	"net/http"
	"storage/internal/api/restapi"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"golang.org/x/net/context"
)

type (
	service struct {
		adminUsername string
		adminPass     string
	}

	// Configuration contains config for api service.
	Configuration struct {
		AdminUsername string
		AdminPass     string

		Port int
		Host string
	}

	// Сtx is a synonym for convenience.
	Сtx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

// Serve must be called once before using this package.
func Serve(log Log, cfg Configuration) error {
	//svc := service{
	//	adminPass:     cfg.AdminPass,
	//	adminUsername: cfg.AdminUsername,
	//}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return errors.Wrapf(err, "failed loads embedded")
	}

	searchAPI := restapi.BuildAPI(
		swaggerSpec,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	server := restapi.NewServer(searchAPI)
	// The middleware executes before anything.
	globalMiddleware := func(handler http.Handler) http.Handler {
		logger := makeLogger(swaggerSpec.BasePath())
		return logger(recovery(handleCORS(handler)))
	}
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	middlewares := func(handler http.Handler) http.Handler {
		return accessLog(handler)
	}

	server.SetHandler(globalMiddleware(searchAPI.Serve(middlewares)))
	server.Port = cfg.Port
	server.Host = cfg.Host

	defer log.WarnIfFail(server.Shutdown)

	log.Info("protocol", "version", swaggerSpec.Spec().Info.Version, "address", net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)))
	return server.Serve()
}

const (
	logUser = "logUser"
	notAuth = 0
)

func fromRequest(r *http.Request, userID int) (Сtx, Log) {
	ctx := r.Context()

	log := structlog.FromContext(ctx, nil).SetDefaultKeyvals(logUser, userID)
	return ctx, log
}
