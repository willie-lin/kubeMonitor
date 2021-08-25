package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sReplicaSet holds the schema definition for the K8sReplicaSet entity.
type K8sReplicaSet struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sReplicaSet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sReplicaSets"},
	}
}

// Mixin Container
func (K8sReplicaSet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sReplicaSet.
func (K8sReplicaSet) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name").MaxLen(256),

		field.Uint("k8sClusterId"),
		field.Uint("k8sNamespaceId"),
		field.Uint("k8sDeploymentId"),
		field.Uint("k8sObjectId"),

	}
}

// Edges of the K8sReplicaSet.
func (K8sReplicaSet) Edges() []ent.Edge {
	return nil
}
