package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Annotations of the Cluster.
func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "events"},
	}
}

// Mixin Container
func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Time("ts"),
		field.Float("value"),

		field.Bool("acked"),
		field.Time("ackedTs"),

		field.Uint("endpointId"),
		field.Uint("typeId"),
		field.Uint("nameId"),
		field.Uint("labelId"),

		field.Uint("clusterId"),
		field.Uint("agentId"),
		field.Uint("nodeId"),
		field.Uint("procesId"),
		field.Uint("containerId"),
		field.Uint("podId"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("MetricName_events", MetricName.Type).Ref("events").Unique(),
		edge.From("MetricLabel_events", MetricLabel.Type).Ref("events").Unique(),
		edge.From("MetricEndpoint_events", MetricEndpoint.Type).Ref("events").Unique(),
	}
}
