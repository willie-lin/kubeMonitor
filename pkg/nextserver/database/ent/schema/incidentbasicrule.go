package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// IncidentBasicRule holds the schema definition for the IncidentBasicRule entity.
type IncidentBasicRule struct {
	ent.Schema
}

// Annotations of the Cluster.
func (IncidentBasicRule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "incidentBasicRules"},
	}
}

// Mixin Container
func (IncidentBasicRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the IncidentBasicRule.
func (IncidentBasicRule) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name"),

		field.String("Description"),
		field.String("query"),
	}
}

// Edges of the IncidentBasicRule.
func (IncidentBasicRule) Edges() []ent.Edge {
	return nil
}
