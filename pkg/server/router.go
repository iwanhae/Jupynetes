package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/iwanhae/Jupynetes/pkg/config"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

var tokenAuth *jwtauth.JWTAuth

//InitRouter Initialize router
func InitRouter(ctx context.Context, c *config.Configs) *chi.Mux {
	r := chi.NewRouter()

	// Default middlwares
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RealIP)

	// Logging
	r.Use(hlog.NewHandler(log.Logger))
	r.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
	r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		log.Ctx(r.Context()).Info().
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
	r.Use(middleware.AllowContentType("application/json"))

	tokenAuth = jwtauth.New("HS256", []byte(c.SecretKey), nil)

	// User Application
	r.Route("/v1", func(r chi.Router) {
		v1 := Server{}

		// Without Authorization
		r.Group(func(r chi.Router) {
			r.Post("/login", v1.LoginUser)
			r.Get("/logout", v1.LogoutUser)
		})

		// User Authorized
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(AuthorizeUser)

			r.Get("/server", v1.GetServerList)
			r.Post("/server", v1.CreateServer)
			r.Delete("/server/{serverId}", v1.DeleteServer)
			r.Get("/server/{serverId}", v1.GetServer)
			r.Get("/template", v1.GetTemplateList)
			r.Get("/user", v1.GetUserInfo)
			r.Post("/user", v1.UpdateUserInfo)

			// Admin Authorized
			r.Group(func(r chi.Router) {
				r.Use(chkAdmin)

				r.Post("/admin/quota", v1.AdminSetQuota)
				r.Post("/admin/quota/{userId}", v1.AdminSetQuota)
				r.Post("/admin/template", v1.AdminCreateTemplate)
				r.Get("/admin/user", v1.AdminGetUserList)
			})
		})
	})
	log.Info().Msg("initailizing router finished")
	return r
}

func chkAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if GetUser(r.Context()) != "admin" {
			send(w, http.StatusUnauthorized, Reason{"you are not admin :-("})
			return
		}
		next.ServeHTTP(w, r)
	})
}
