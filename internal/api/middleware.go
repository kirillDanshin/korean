package api

import (
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/powerman/structlog"
	"github.com/rs/cors"
	"github.com/sebest/xff"
)

type middlewareFunc func(http.Handler) http.Handler

const (
	logHTTPMethod = "httpMethod"
	logHTTPStatus = "httpStatus"
	logFunc       = "func"
	logRemote     = "remote"
)

func makeLogger(basePath string) middlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			remote := xff.GetRemoteAddr(r)
			log := structlog.New(
				logRemote, remote,
				logHTTPMethod, r.Method,
				logFunc, strings.TrimPrefix(r.URL.Path, basePath),
			)
			r = r.WithContext(structlog.NewContext(r.Context(), log))

			handler.ServeHTTP(w, r)
		})
	}
}

func recovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			switch err := recover(); err := err.(type) {
			default:
				log := structlog.FromContext(r.Context(), nil)
				log.PrintErr(err, structlog.KeyStack, structlog.Auto)
				os.Exit(2)
			case nil:
			case net.Error:
				log := structlog.FromContext(r.Context(), nil)
				log.PrintErr(err)
			}
		}()

		handler.ServeHTTP(w, r)
	})
}

func handleCORS(next http.Handler) http.Handler {
	return cors.AllowAll().Handler(next)
}

func accessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		sw := statusWriter{ResponseWriter: w}
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log := structlog.FromContext(r.Context(), nil)
		if code := sw.status; code < 500 {
			log.Info("handled", "in", duration, logHTTPStatus, code)
		} else {
			log.PrintErr("failed to handle", "in", time.Since(start), logHTTPStatus, code)
		}
	})
}
