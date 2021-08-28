package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// MetricName holds the schema definition for the MetricName entity.
type MetricName struct {
	ent.Schema
}

// Annotations of the Cluster.
func (MetricName) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "metricNames"},
	}
}

// Mixin Container
func (MetricName) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the MetricName.
func (MetricName) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),
		field.String("name").MaxLen(256).Unique(),
		field.String("help"),

		field.Uint("typeId"),
	}
}

// Edges of the MetricName.
func (MetricName) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metrics", Metric.Type),
		edge.To("events", Event.Type),
		edge.From("MetricType_MetricNames", MetricType.Type).Ref("MetricNames").Unique(),
	}
}

func (MetricName) indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "typeId"),
	}
}
