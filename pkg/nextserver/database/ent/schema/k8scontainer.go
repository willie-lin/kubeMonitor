package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sContainer holds the schema definition for the K8sContainer entity.
type K8sContainer struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sContainer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sContainers"},
	}
}

// Mixin Container
func (K8sContainer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the K8sContainer.
func (K8sContainer) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name").MaxLen(256),
		field.String("image").MaxLen(256),
		field.String("containerType").MaxLen(64),
		field.String("containerId").MaxLen(256),

		field.Uint("k8sClusterId"),
		field.Uint("k8sNamespaceId"),
		field.Uint("K8sPodId"),

	}
}

// Edges of the K8sContainer.
func (K8sContainer) Edges() []ent.Edge {
	return nil
}
