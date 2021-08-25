package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Proces holds the schema definition for the Proces entity.
type Proces struct {
	ent.Schema
}

// Annotations of the Cluster.
func (Proces) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "process"},
	}
}

// Fields of the Process.
func (Proces) Fields() []ent.Field {
	return []ent.Field {
		field.Uint("id").Unique(),
		field.String("name"),
		field.Int32("pId"),
		field.String("cmd"),
		field.JSON("info", []string{}),
		field.String("clusterId"),
		field.String("nodeId"),
		field.String("containerId"),
	}
}

// Edges of the Process.
func (Proces) Edges() []ent.Edge {
	return nil
}

// Mixin Process
func (Proces) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Proces) Indexes() []ent.Index  {
	return []ent.Index {
		index.Fields("clusterId",  "nodeId", "containerId"),
	}
}