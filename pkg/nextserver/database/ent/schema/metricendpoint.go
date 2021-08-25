package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// MetricEndpoint holds the schema definition for the MetricEndpoint entity.
type MetricEndpoint struct {
	ent.Schema
}



// Annotations of the Cluster.
func (MetricEndpoint) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "metricEndpoints"},
	}
}

// Mixin Container
func (MetricEndpoint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}



// Fields of the MetricEndpoint.
func (MetricEndpoint) Fields() []ent.Field {
	return []ent.Field {
		field.Uint("id").Unique(),
		field.String("path").MaxLen(256).Unique(),
	}
}

// Edges of the MetricEndpoint.
func (MetricEndpoint) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("metrics", Metric.Type),
		edge.To("events", Event.Type),
	}
}

func (MetricEndpoint) Indexes() []ent.Index  {
	return []ent.Index {
		index.Fields("path"),
	}
}
