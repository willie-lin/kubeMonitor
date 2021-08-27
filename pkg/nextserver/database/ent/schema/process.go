package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Process holds the schema definition for the Process entity.
type Process struct {
	ent.Schema
}

// Annotations of the Cluster.
func (Process) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "processes"},
	}
}

// Fields of the Process.
func (Process) Fields() []ent.Field {
	return []ent.Field{
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
func (Process) Edges() []ent.Edge {
	return nil
}

// Mixin Process
func (Process) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Process) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("clusterId", "nodeId", "containerId"),
	}
}
