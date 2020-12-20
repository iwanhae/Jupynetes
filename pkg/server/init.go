package server

import (
	"net/http"
	"time"

	"git.iwanhae.kr/wan/jupynetes/pkg/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

//InitRouter Initialize router
func InitRouter(c config.Configs) *chi.Mux {
	r := chi.NewRouter()

	// Default middlwares
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RealIP)

	// Logging
	r.Use(hlog.NewHandler(log.Logger))
	r.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
	r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("user_agent", r.Header.Get("User-Agent")).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))

	// Panic Recover
	r.Use(middleware.Recoverer)

	// User Application

	log.Info().Msg("initailizing router finished")
	return r
}
