package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/uuid"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("domain").Unique(),
		field.String("callback"),
		field.Strings("tags"),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("surveys", Survey.Type),
		edge.To("users", Person.Type),
		edge.To("admins", Person.Type),
	}
}
