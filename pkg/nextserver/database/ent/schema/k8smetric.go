package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// K8sMetric holds the schema definition for the K8sMetric entity.
type K8sMetric struct {
	ent.Schema
}



// Annotations of the Cluster.
func (K8sMetric) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sMetrics"},
	}
}

// Mixin Container
func (K8sMetric) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sMetric.
func (K8sMetric) Fields() []ent.Field {
	return []ent.Field {
		field.Time("ts"),
		field.Float("value"),

		field.Uint("endpointId"),
		field.Uint("typeId"),
		field.Uint("nameId"),
		field.Uint("labelId"),

		field.Uint("k8sClusterId"),
		field.Uint("k8sNodeId"),
		field.Uint("k8sNamespaceId"),
		field.Uint("k8sContainerId"),

	}
}

// Edges of the K8sMetric.
func (K8sMetric) Edges() []ent.Edge {
	return nil
}
