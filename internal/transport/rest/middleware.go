package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ysle0/go-starter/internal/cfg"
	"net/http"
	"time"
)

func useMiddlewares(r *chi.Mux, cfg cfg.Config) {
	// API timeout
	timeoutDuration := cfg.Server.HandlerResponseTimeoutInMs
	hRespTimeoutDuration := time.Duration(timeoutDuration) * time.Millisecond

	middlewares := []func(http.Handler) http.Handler{
		middleware.Timeout(hRespTimeoutDuration),
		middleware.Recoverer,
		middleware.Logger,
		middleware.RequestID,
		middleware.RealIP,
		middleware.StripSlashes,
		middleware.RedirectSlashes,
		middleware.Heartbeat("/ping"),
		middleware.Compress(3, "gzip"),
	}

	for _, mw := range middlewares {
		r.Use(mw)
	}
}
