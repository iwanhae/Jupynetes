package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/iwanhae/Jupynetes/pkg/config"
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

	server := Server{}
	// User Application

	// Without Authorization
	r.Group(func(r chi.Router) {
		r.Post("/login", server.LoginUser)
		r.Get("/logout", server.LogoutUser)
	})

	// User Authorized
	r.Group(func(r chi.Router) {
		r.Get("/server", server.GetServerList)
		r.Post("/server", server.CreateServer)
		r.Delete("/server/{serverId}", server.DeleteServer)
		r.Get("/server/{serverId}", server.GetServer)
		r.Get("/template", server.GetTemplateList)
		r.Get("/user", server.GetUserInfo)
		r.Post("/user", server.UpdateUserInfo)
	})

	// Admin Authorized
	r.Group(func(r chi.Router) {
		r.Post("/admin/quota", server.AdminSetQuota)
		r.Post("/admin/quota/{userId}", server.AdminSetQuota)
		r.Post("/admin/template", server.AdminCreateTemplate)
		r.Get("/admin/user", server.AdminGetUserList)
	})
	log.Info().Msg("initailizing router finished")
	return r
}
