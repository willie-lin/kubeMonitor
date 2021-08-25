package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sEvent holds the schema definition for the K8sEvent entity.
type K8sEvent struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sEvents"},
	}
}

// Mixin Container
func (K8sEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sEvent.
func (K8sEvent) Fields() []ent.Field {
	return []ent.Field {
		field.Time("ts"),
		field.Float("value"),

		field.Uint("endpointId"),
		field.Uint("typeId"),
		field.Uint("nameId"),
		field.Uint("labelId"),

		field.Uint("clusterId"),
		field.Uint("namespaceId"),
		field.Uint("nodeId"),
		field.Uint("containerId"),
		field.Uint("podId"),
	}
}

// Edges of the K8sEvent.
func (K8sEvent) Edges() []ent.Edge {
	return nil
}
