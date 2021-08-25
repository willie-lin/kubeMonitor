package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sNamespace holds the schema definition for the K8sNamespace entity.
type K8sNamespace struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sNamespace) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sNamespaces"},
	}
}

// Mixin Container
func (K8sNamespace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sNamespace.
func (K8sNamespace) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name").MaxLen(128),

		field.Uint("k8sClusterId"),
		field.Uint("k8sObjectId"),

	}
}

// Edges of the K8sNamespace.
func (K8sNamespace) Edges() []ent.Edge {
	return nil
}
