package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/iwanhae/Jupynetes/pkg/common"
)

// Template holds the schema definition for the Template entity.
type Template struct {
	ent.Schema
}

// Fields of the Template.
func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.String("template").SchemaType(map[string]string{
			dialect.MySQL: "longtext",
		}),
		field.JSON("variables", &common.TemplateVariables{}),
		field.Time("created_at").Default(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Template.
func (Template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("templates"),
		edge.From("server", Server.Type).Ref("template_from"),
	}
}
