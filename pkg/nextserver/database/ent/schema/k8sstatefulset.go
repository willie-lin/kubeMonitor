package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sStatefulSet holds the schema definition for the K8sStatefulSet entity.
type K8sStatefulSet struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sStatefulSet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sStatefulSets"},
	}
}

// Mixin Container
func (K8sStatefulSet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sStatefulSet.
func (K8sStatefulSet) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name").MaxLen(256),

		field.Uint("k8sClusterId"),
		field.Uint("k8sNamespaceId"),
		field.Uint("k8sObjectId"),

	}

}

// Edges of the K8sStatefulSet.
func (K8sStatefulSet) Edges() []ent.Edge {
	return nil
}
