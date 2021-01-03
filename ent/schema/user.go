package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Unique(),
		field.String("user_pw"),
		field.Int("quota_instance").Default(0),
		field.Int("quota_cpu").Default(0),
		field.Int("quota_memory").Default(0),
		field.Int("quota_nvidia_gpu").Default(0),
		field.Int("quota_storage").Default(0),
		field.Time("created_at").Default(time.Now),
		field.Time("deleted_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("user"),

		edge.To("servers", Server.Type),
		edge.To("templates", Template.Type),
	}
}
