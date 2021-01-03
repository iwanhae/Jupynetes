package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/iwanhae/Jupynetes/pkg/kubeclient"
)

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("template"),
		field.JSON("variables", &[]kubeclient.Variable{}),
		field.String("ip"),
		field.String("description"),
		field.Int("cpu"),
		field.Int("memory"),
		field.Int("nvidia_gpu"),
		field.Time("created_at").Default(time.Now),
		field.Time("deleted_at"),
	}
}

// Edges of the Server.
func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owners", User.Type).Ref("servers"),

		edge.To("template_from", Template.Type),
	}
}
