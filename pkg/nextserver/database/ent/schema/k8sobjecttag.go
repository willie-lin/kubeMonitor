package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sObjectTag holds the schema definition for the K8sObjectTag entity.
type K8sObjectTag struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sObjectTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sObjectTags"},
	}
}

// Mixin Container
func (K8sObjectTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sObjectTag.
func (K8sObjectTag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.Uint("k8sObjectId"),
		field.Uint("k8sLabelId"),

	}
}

// Edges of the K8sObjectTag.
func (K8sObjectTag) Edges() []ent.Edge {
	return nil
}
