package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cluster holds the schema definition for the Cluster entity.
type Cluster struct {
	ent.Schema
}

// Annotations of the Cluster.
func (Cluster) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "clusters"},
	}
}

// Fields of the Cluster.
func (Cluster) Fields() []ent.Field {

	return []ent.Field{
		field.Uint("id").Unique(),
		field.String("name").MaxLen(128),
		field.String("description"),
		field.Bool("disabled"),

	}
}

// Edges of the Cluster.
func (Cluster) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("agents", Agent.Type),
		edge.To("nodes", Node.Type),
	}
}

// Mixin Cluster
func (Cluster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
