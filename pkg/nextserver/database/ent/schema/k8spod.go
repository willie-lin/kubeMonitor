package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sPod holds the schema definition for the K8sPod entity.
type K8sPod struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sPod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sPods"},
	}
}

// Mixin Container
func (K8sPod) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the K8sPod.
func (K8sPod) Fields() []ent.Field {

	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("name").MaxLen(256),
		field.String("qos").MaxLen(32),

		field.Uint("k8sClusterId"),
		field.Uint("k8sNamespaceId"),
		field.Uint("k8sObjectId"),

	}

}

// Edges of the K8sPod.
func (K8sPod) Edges() []ent.Edge {
	return nil
}
