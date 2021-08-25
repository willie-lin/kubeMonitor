package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Metric holds the schema definition for the Metric entity.
type Metric struct {
	ent.Schema
}


// Annotations of the Cluster.
func (Metric) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "metrics"},
	}
}

// Mixin Container
func (Metric) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Metric.
func (Metric) Fields() []ent.Field {
	return []ent.Field {
		field.Time("ts"),
		field.Float("value"),

		field.Uint("endpointId"),
		field.Uint("typeId"),
		field.Uint("nameId"),
		field.Uint("labelId"),

		field.Uint("clusterId"),
		field.Uint("nodeId"),
		field.Uint("procesId"),
		field.Uint("containerId"),
	}
}

// Edges of the Metric.
func (Metric) Edges() []ent.Edge {
	return nil
}


func (Metric) Indexes() []ent.Index  {
	return []ent.Index {
		index.Fields("ts", "value", "endpointId",
			"typeId", "nameId", "labelId", "clusterId",
			"nodeId", "procesId", "containerId"),
	}
}
