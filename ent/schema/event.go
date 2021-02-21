package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("message").SchemaType(map[string]string{
			dialect.MySQL: "longtext",
		}),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
		edge.To("server", Server.Type),
	}
}
