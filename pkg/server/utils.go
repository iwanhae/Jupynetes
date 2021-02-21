package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/iwanhae/Jupynetes/ent"
	"github.com/iwanhae/Jupynetes/ent/event"
	"github.com/iwanhae/Jupynetes/pkg/common"
	"github.com/iwanhae/Jupynetes/pkg/database"
)

func CreateEvent(ctx context.Context, msg string, s *ent.Server, u *ent.User) error {
	_, err := database.GetClient().Event.Create().
		SetMessage(msg).
		AddUser(u).
		AddServer(s).
		Save(ctx)
	return err
}

//GetServerObject Get Server Object by querying db
func GetServerObject(ctx context.Context, s *ent.Server) (*common.ServerObject, error) {
	res := &common.ServerObject{
		Id: s.ID,

		Name:        s.Name,
		Description: s.Description,
		Flavor: common.Flavor{
			Cpu:       s.CPU,
			Memory:    s.Memory,
			NvidiaGpu: s.NvidiaGpu,
		},
		CreatedAt: s.CreatedAt,
	}

	owners, err := s.QueryOwners().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, o := range owners {
		res.Owner = append(res.Owner, o.UserID)
	}

	template, err := s.QueryTemplateFrom().First(ctx)
	if err != nil {
		return nil, err
	}

	res.Template = *template.ToCommonType()
	res.Template.Variables = *s.Variables

	e, err := s.QueryEvent().Order(ent.Desc(event.FieldCreatedAt)).First(ctx)
	if err != nil {
		return nil, err
	}

	res.Message = e.Message
	return res, nil
}

func send(w http.ResponseWriter, statusCode int, data interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(buf.Bytes())
}
