package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// K8sCluster holds the schema definition for the K8sCluster entity.
type K8sCluster struct {
	ent.Schema
}

// Annotations of the Cluster.
func (K8sCluster) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sClusters"},
	}
}

// Mixin Container
func (K8sCluster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the K8sCluster.
func (K8sCluster) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),
		field.String("name").MaxLen(128),
		field.Uint("agentClusterId"),


	}
}

// Edges of the K8sCluster.
func (K8sCluster) Edges() []ent.Edge {
	return nil
}

func (K8sCluster) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
