package server

import "net/http"

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
func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {}

// LogoutUser Log out user by clear cookie
// (GET /logout)
func (s *Server) LogoutUser(w http.ResponseWriter, r *http.Request) {}

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
