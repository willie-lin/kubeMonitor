package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sDaemonSet holds the schema definition for the K8sDaemonSet entity.
type K8sDaemonSet struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sDaemonSet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sDaemonSets"},
	}
}

// Mixin Container
func (K8sDaemonSet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}




// Fields of the K8sDaemonSet.
func (K8sDaemonSet) Fields() []ent.Field {

	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name").MaxLen(256),

		field.Uint("k8sClusterId"),
		field.Uint("k8sNamespaceId"),
		field.Uint("k8sObjectId"),
	}
}

// Edges of the K8sDaemonSet.
func (K8sDaemonSet) Edges() []ent.Edge {
	return nil
}
