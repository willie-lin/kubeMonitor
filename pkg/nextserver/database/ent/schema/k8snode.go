package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sNode holds the schema definition for the K8sNode entity.
type K8sNode struct {
	ent.Schema
}

// Annotations of the Cluster.
func (K8sNode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sNodes"},
	}
}

// Mixin Container
func (K8sNode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sNode.
func (K8sNode) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name").MaxLen(128),

		field.Uint("k8sClusterId"),
		field.Uint("k8sObjectId"),

	}
}

// Edges of the K8sNode.
func (K8sNode) Edges() []ent.Edge {
	return nil
}
