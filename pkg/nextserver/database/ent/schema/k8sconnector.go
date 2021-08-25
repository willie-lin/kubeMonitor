package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// K8sConnector holds the schema definition for the K8sConnector entity.
type K8sConnector struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sConnector) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sConnectors"},
	}
}

// Mixin Container
func (K8sConnector) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sConnector.
func (K8sConnector) Fields() []ent.Field {
	return []ent.Field {
		field.Uint("id").Unique(),

		field.String("name").MaxLen(128),
		field.String("status").MaxLen(32),
		field.String("method").MaxLen(32),

		field.Bool("inCluster"),
		field.String("bearerToken"),
		field.String("kubeConfig"),

	}
}

// Edges of the K8sConnector.
func (K8sConnector) Edges() []ent.Edge {
	return nil
}

func (K8sConnector) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

