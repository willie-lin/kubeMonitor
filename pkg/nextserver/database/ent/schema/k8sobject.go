package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sObject holds the schema definition for the K8sObject entity.
type K8sObject struct {
	ent.Schema
}

// Annotations of the Cluster.
func (K8sObject) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sObjects"},
	}
}

// Mixin Container
func (K8sObject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sObject.
func (K8sObject) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.Uint("k8sClusterId"),

		field.String("apiVersion").MaxLen(128),
		field.String("kind").MaxLen(128),
		field.String("name").MaxLen(256),

		field.JSON("metadata", []string{}),
		field.JSON("spec", []string{}),
		field.JSON("status", []string{}),


	}
}

// Edges of the K8sObject.
func (K8sObject) Edges() []ent.Edge {
	return nil
}
