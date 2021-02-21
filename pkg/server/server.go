package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/iwanhae/Jupynetes/ent/server"
	"github.com/iwanhae/Jupynetes/ent/template"
	"github.com/iwanhae/Jupynetes/ent/user"
	"github.com/iwanhae/Jupynetes/pkg/database"
	"github.com/iwanhae/Jupynetes/pkg/kubeclient"
	"github.com/rs/zerolog/log"

	"github.com/iwanhae/Jupynetes/pkg/common"
)

//Server dummy routing sturct
type Server struct{}

// AdminSetQuota set user quota
// (POST /admin/quota) and (POST /admin/quota/{userId})
func (s *Server) AdminSetQuota(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &common.Quota{}
	if err := render.Bind(r, req); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("invalid format")
		send(w, http.StatusBadRequest, common.Reason{
			Reason: "Invalid format",
		})
		return
	}
	db := database.GetClient()

	userID := chi.URLParam(r, "userId")
	if len(userID) == 0 {
		send(w, http.StatusNotFound, common.GetReasonf("Set global quota is not supported by api"))
		return
	}
	u, err := db.User.Query().Where(user.UserIDEQ(userID)).First(ctx)
	if err != nil {
		send(w, http.StatusNotFound, common.GetReasonf("user not found:%s", userID))
		return
	}

	u, err = db.User.UpdateOneID(u.ID).
		SetQuotaCPU(req.Cpu).
		SetQuotaInstance(req.Instance).
		SetQuotaMemory(req.Memory).
		SetQuotaStorage(req.Storage).
		SetQuotaNvidiaGpu(req.NvidiaGpu).
		Save(ctx)

	if err != nil {
		send(w, http.StatusInternalServerError, common.GetReason(err.Error()))
		return
	}
	send(w, http.StatusAccepted, req)
	return
}

// AdminCreateTemplate create template
// (POST /admin/template)
func (s *Server) AdminCreateTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &common.Template{}
	if err := render.Bind(r, req); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("invalid format")
		send(w, http.StatusBadRequest, common.Reason{
			Reason: "Invalid format",
		})
		return
	}
	log.Ctx(ctx).Info().Interface("template", req).Msg("will create template")

	db := database.GetClient()
	templates, err := db.Template.Query().Where(
		template.NameEQ(req.Name),
	).All(ctx)

	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("fail to query template from db")
		send(w, http.StatusInternalServerError, common.GetReason("unexpected db error"))
		return
	}

	if len(templates) != 0 {
		log.Ctx(ctx).Error().Str("template_name", req.Name).Msg("template name already exists")
		send(w, http.StatusConflict, common.GetReasonf("%q already exists", req.Name))
		return
	}

	// TODO: Template Validation with kubectl --dry-run

	template, err := db.Template.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetTemplate(req.Body).
		SetVariables(&req.Variables).Save(ctx)

	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("fail to query template to db")
		send(w, http.StatusInternalServerError, common.GetReason("unexpected db error"))
		return
	}

	res := template.ToCommonType()

	send(w, http.StatusAccepted, res)
}

// AdminGetUserList get user list
// (GET /admin/user)
func (s *Server) AdminGetUserList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	db := database.GetClient()
	users, err := db.User.Query().All(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("fail to query user list from db")
		send(w, http.StatusInternalServerError, common.GetReason("unexpected db error"))
	}

	res := []common.UserInfo{}
	for _, user := range users {
		res = append(res, common.UserInfo{
			Id: user.UserID,
			UserQuota: common.Quota{
				Instance:  user.QuotaInstance,
				Cpu:       user.QuotaCPU,
				Memory:    user.QuotaMemory,
				NvidiaGpu: user.QuotaNvidiaGpu,
				Storage:   user.QuotaStorage,
			},
			GroupQuota: rootQuota,
		})
	}
	send(w, http.StatusOK, res)
}

// LoginUser Logs in user by set cookie
// (POST /login)
func (s *Server) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &common.LoginRequest{}
	if err := render.Bind(r, req); err != nil {
		send(w, http.StatusBadRequest, common.Reason{
			Reason: "Invalid format",
		})
		return
	}

	db := database.GetClient()
	user, err := db.User.Query().Where(user.UserIDEQ(req.Id)).All(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("fail to query db")
		send(w, http.StatusInternalServerError, common.GetReason("unexpected db error"))
		return
	}
	if len(user) == 0 {
		log.Ctx(ctx).Error().
			Str("user_id", req.Id).
			Msg("auth failed:user not found")
		send(w, http.StatusBadRequest, common.GetReason("invalid id or pw"))
		return
	}
	if database.IsEqualPassword(user[0].UserPw, req.Pw) == false {
		log.Ctx(ctx).Error().
			Str("user_id", req.Id).
			Msg("auth failed:wrong password")
		send(w, http.StatusBadRequest, common.GetReason("invalid id or pw"))
		return
	}
	cookie, err := GenerateTokenCookie(req.Id)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Str("user_id", req.Id).Msg("fail to write jwt")
		send(w, http.StatusInternalServerError, common.GetReason("signing error"))
		return
	}
	http.SetCookie(w, cookie)
	send(w, http.StatusOK, common.GetReasonf("Welcome %q :-)", req.Id))

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
	send(w, http.StatusOK, common.GetReason("you are logged out. :-)"))
	return
}

// GetServerList Get list of accessible server to user
// (GET /server)
func (s *Server) GetServerList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := GetUser(ctx)
	db := database.GetClient()
	u, err := db.User.Query().Where(user.UserID(userID)).First(ctx)
	if err != nil {
		send(w, http.StatusForbidden, common.GetReasonf("unkown user error:%s", err.Error()))
		return
	}
	servers, err := u.QueryServers().All(ctx)

	res := []common.ServerObject{}
	for _, s := range servers {
		server, err := GetServerObject(ctx, s)
		if err != nil {
			send(w, http.StatusInternalServerError, common.GetReasonf("%s", err.Error()))
			return
		}
		res = append(res, *server)
	}
	send(w, http.StatusOK, res)
	return
}

// CreateServer Create server request
// (POST /server)
func (s *Server) CreateServer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &common.CreateServerRequest{}
	if err := render.Bind(r, req); err != nil {
		send(w, http.StatusBadRequest, common.GetReasonf("invalid format:%s", err.Error()))
		return
	}

	db := database.GetClient()

	user, err := db.User.Query().Where(user.UserIDEQ(GetUser(ctx))).First(ctx)
	if err != nil {
		send(w, http.StatusBadRequest, common.GetReason("user info not found"))
		return
	}

	servers, err := db.Server.Query().Where(server.NameEQ(req.Name)).All(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("fail to query db")
		send(w, http.StatusInternalServerError, common.GetReason("unexpected db error"))
		return
	}
	if len(servers) != 0 {
		log.Ctx(ctx).Error().Interface("server", servers[0]).Msg("fail to create server:duplicate name")
		send(w, http.StatusConflict, common.GetReason("duplicated name"))
		return
	}

	templateEnt, err := db.Template.Get(ctx, req.TemplateId)
	if err != nil {
		send(w, http.StatusBadRequest, common.GetReasonf("Given template id not exists:%s", err.Error()))
		return
	}

	template := templateEnt.ToCommonType()

	//will apply variable from first to last in sequence
	template.Variables = append(template.Variables /* default */, req.TemplateVariables... /* user custommed */)

	//embedded flavor info to TemplateVariable
	if err := req.Flavor.Bind(r); err == nil {
		template.Variables = append(template.Variables,
			common.TemplateVariable{
				Name:  kubeclient.VariableFlavorCPU,
				Value: fmt.Sprintf("%d", req.Flavor.Cpu),
			},
			common.TemplateVariable{
				Name:  kubeclient.VariableFlavorMem,
				Value: fmt.Sprintf("%d", req.Flavor.Memory),
			},
			common.TemplateVariable{
				Name:  kubeclient.VariableFlavorNvidiaGpu,
				Value: fmt.Sprintf("%d", req.Flavor.NvidiaGpu),
			},
		)
	}

	created, err := db.Server.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetTemplate(template.Body).
		SetIP("0.0.0.0").
		SetVariables(&template.Variables).
		SetCPU(req.Flavor.Cpu).
		SetMemory(req.Flavor.Memory).
		SetNvidiaGpu(req.Flavor.NvidiaGpu).
		AddOwners(user).
		AddTemplateFrom(templateEnt).
		Save(ctx)

	if err != nil {
		send(w, http.StatusInternalServerError, common.GetReasonf("fail to create server:fail to update db:%s", err.Error()))
		return
	}

	CreateEvent(ctx, fmt.Sprintf("creating server"), created, user)

	err = kubeclient.DeployServer(ctx, req.Name, template)
	if err != nil {
		CreateEvent(ctx, fmt.Sprintf("fail to create server:%s", err.Error()), created, user)
		send(w, http.StatusInternalServerError, common.GetReasonf("fail to create server:fail to apply to kubernetes:%s", err.Error()))
		return
	}
	CreateEvent(ctx, fmt.Sprintf("request accepted"), created, user)

	res, err := GetServerObject(ctx, created)
	if err != nil {
		send(w, http.StatusInternalServerError, common.GetReasonf("fail to get server info", err.Error()))
		return
	}

	send(w, http.StatusAccepted, res)
	return
}

// DeleteServer Delete server
// (DELETE /server/{serverId}
func (s *Server) DeleteServer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	serverID, err := strconv.Atoi(chi.URLParam(r, "serverId"))
	if err != nil {
		send(w, http.StatusBadRequest, common.GetReason("can not parse server id"))
		return
	}

	db := database.GetClient()
	server, err := db.Server.Query().Where(server.ID(serverID)).First(ctx)
	if err != nil || server == nil {
		send(w, http.StatusNotFound, common.GetReasonf("server not found:%s", err.Error()))
		return
	}

	tmpErr := kubeclient.DeleteServer(ctx, server.Name)

	user, _ := db.User.Query().Where(user.UserID(GetUser(ctx))).First(ctx)
	CreateEvent(ctx, "server deleted", server, user)

	err = db.Server.DeleteOneID(server.ID).Exec(ctx)
	if err != nil {
		send(w, http.StatusInternalServerError, common.GetReasonf("fail to delete server:%s", err.Error()))
		return
	}
	if tmpErr != nil {
		send(w, http.StatusAccepted, common.GetReasonf("server deleted, but warnings:%s", tmpErr))
		return
	}
	send(w, http.StatusAccepted, common.GetReasonf("server deleted"))
	return
}

// GetServer Get server info
// (GET /server/{serverId}
func (s *Server) GetServer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	serverID, err := strconv.Atoi(chi.URLParam(r, "serverId"))
	if err != nil {
		send(w, http.StatusBadRequest, common.GetReason("can not parse server id"))
		return
	}

	db := database.GetClient()
	server, err := db.Server.Query().Where(server.ID(serverID)).First(ctx)
	if err != nil {
		send(w, http.StatusNotFound, common.GetReasonf("fail to find server %q", serverID))
		return
	}

	res, err := GetServerObject(ctx, server)
	if err != nil {
		send(w, http.StatusInternalServerError, common.GetReasonf("fail to get server %q", serverID))
		return
	}

	userID := GetUser(ctx)
	if userID == "admin" {
		send(w, http.StatusOK, res)
		return
	}
	for _, o := range res.Owner {
		if o == userID {
			send(w, http.StatusOK, res)
			return
		}
	}
	send(w, http.StatusUnauthorized, common.GetReason("you do not have a permission to access this server"))
}

// GetTemplateList get template list
// (GET /template)
func (s *Server) GetTemplateList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	db := database.GetClient()
	templates, err := db.Template.Query().All(ctx)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("fail to query template list from db")
		send(w, http.StatusInternalServerError, common.GetReason("unexpected db error"))
	}

	var resp []*common.Template

	for _, t := range templates {
		resp = append(resp, t.ToCommonType())
	}

	send(w, http.StatusOK, resp)
	return
}

// GetUserInfo get user info
// (GET /user)
func (s *Server) GetUserInfo(w http.ResponseWriter, r *http.Request) {}

// UpdateUserInfo update user info
// (POST /user)
func (s *Server) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {}
