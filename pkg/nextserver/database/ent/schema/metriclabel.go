package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// MetricLabel holds the schema definition for the MetricLabel entity.
type MetricLabel struct {
	ent.Schema
}

// Annotations of the Cluster.
func (MetricLabel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "metricLabels"},
	}
}

// Mixin Container
func (MetricLabel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}



// Fields of the MetricLabel.
func (MetricLabel) Fields() []ent.Field {
	return []ent.Field {
		field.Uint("id").Unique(),
		field.String("label").MaxLen(256).Unique(),
	}
}

// Edges of the MetricLabel.
func (MetricLabel) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("metrics", Metric.Type),
		edge.To("events", Event.Type),
	}
}

func (MetricLabel) indexes()  []ent.Index {
	return []ent.Index {
		index.Fields("label"),
	}
}

