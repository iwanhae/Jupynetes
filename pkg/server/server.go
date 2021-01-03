package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/iwanhae/Jupynetes/ent/user"
	"github.com/iwanhae/Jupynetes/pkg/database"
	"github.com/rs/zerolog/log"
)

//Server dummy routing sturct
type Server struct{}

// AdminSetQuota set user quota
// (POST /admin/quota) and (POST /admin/quota/{userId})
func (s *Server) AdminSetQuota(w http.ResponseWriter, r *http.Request) {
}

// AdminCreateTemplate create template
// (POST /admin/template)
func (s *Server) AdminCreateTemplate(w http.ResponseWriter, r *http.Request) {

}

// AdminGetUserList get user list
// (GET /admin/user)
func (s *Server) AdminGetUserList(w http.ResponseWriter, r *http.Request) {}

// LoginUser Logs in user by set cookie
// (POST /login)
func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &LoginRequest{}
	if err := render.Bind(r, req); err != nil {
		send(w, http.StatusBadRequest, Reason{
			Reason: "Invalid format",
		})
		return
	}

	db := database.GetClient()
	user, err := db.User.Query().Where(user.UserIDEQ(req.Id)).All(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("fail to query db")
		send(w, http.StatusInternalServerError, Reason{"unexpected db error"})
		return
	}
	if len(user) == 0 {
		log.Ctx(ctx).Error().
			Str("user_id", req.Id).
			Msg("auth failed:user not found")
		send(w, http.StatusBadRequest, Reason{"invalid id or pw"})
		return
	}
	if database.IsEqualPassword(user[0].UserPw, req.Pw) == false {
		log.Ctx(ctx).Error().
			Str("user_id", req.Id).
			Msg("auth failed:wrong password")
		send(w, http.StatusBadRequest, Reason{"invalid id or pw"})
		return
	}
	cookie, err := GenerateTokenCookie(req.Id)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Str("user_id", req.Id).Msg("fail to write jwt")
		send(w, http.StatusInternalServerError, Reason{"signing error"})
		return
	}
	http.SetCookie(w, cookie)
	send(w, http.StatusOK, Reason{fmt.Sprintf("Welcome %q :-)", req.Id)})

	log.Ctx(ctx).Info().Str("user_id", req.Id).Msg("logged in")
	return
}

// LogoutUser Log out user by clear cookie
// (GET /logout)
func (s *Server) LogoutUser(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "jwt",
		Value:  "",
		MaxAge: -1,
	})
	send(w, http.StatusOK, Reason{"you are logged out. :-)"})
	return
}

// GetServerList Get list of accessible server to user
// (GET /server)
func (s *Server) GetServerList(w http.ResponseWriter, r *http.Request) {}

// CreateServer Create server request
// (POST /server)
func (s *Server) CreateServer(w http.ResponseWriter, r *http.Request) {}

// DeleteServer Delete server
// (DELETE /server/{serverId}
func (s *Server) DeleteServer(w http.ResponseWriter, r *http.Request) {}

// GetServer Get server info
// (GET /server/{serverId}
func (s *Server) GetServer(w http.ResponseWriter, r *http.Request) {}

// GetTemplateList get template list
// (GET /template)
func (s *Server) GetTemplateList(w http.ResponseWriter, r *http.Request) {}

// GetUserInfo get user info
// (GET /user)
func (s *Server) GetUserInfo(w http.ResponseWriter, r *http.Request) {}

// UpdateUserInfo update user info
// (POST /user)
func (s *Server) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {}
