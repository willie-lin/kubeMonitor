package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sDeployment holds the schema definition for the K8sDeployment entity.
type K8sDeployment struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sDeployment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sDeployments"},
	}
}

// Mixin Container
func (K8sDeployment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sDeployment.
func (K8sDeployment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name").MaxLen(256),

		field.Uint("k8sClusterId"),
		field.Uint("k8sNamespaceId"),
		field.Uint("k8sObjectId"),

	}
}

// Edges of the K8sDeployment.
func (K8sDeployment) Edges() []ent.Edge {
	return nil
}
