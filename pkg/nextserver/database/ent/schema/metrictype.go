package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// MetricType holds the schema definition for the MetricType entity.
type MetricType struct {
	ent.Schema
}

// Annotations of the Cluster.
func (MetricType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "metricTypes"},
	}
}

// Mixin Container
func (MetricType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the MetricType.
func (MetricType) Fields() []ent.Field {
	return []ent.Field {
		field.Uint("id").Unique(),
		field.String("name").Unique(),

	}
}

// Edges of the MetricType.
func (MetricType) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("metricNames", MetricName.Type),
	}
}

func (MetricType) indexes()  []ent.Index {
	return []ent.Index {
		index.Fields("name"),
	}
}
