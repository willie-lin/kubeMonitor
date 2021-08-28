package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Annotations of the Node.
func (Node) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "nodes"},
	}
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("host").MaxLen(128),
		field.String("ipv4").MaxLen(16),
		field.String("ipv6").MaxLen(40),
		field.String("public_ipv4").MaxLen(16),
		field.String("public_ipv6").MaxLen(40),

		field.String("os").MaxLen(64),

		field.String("platform").MaxLen(64),
		field.String("platformFamily").MaxLen(64),
		field.JSON("info", []string{}),

		field.String("uuid").Unique(),
		field.String("description"),
		field.Bool("disabled"),

		field.Uint("agentId"),
		field.Uint("clusterId"),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("owner", Cluster.Type).Ref("nodes").Unique(),
		edge.To("containers", Container.Type),
		edge.To("process", Proces.Type),
	}
}

// Mixin Node
func (Node) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Node) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("agentId", "clusterId"),
	}
}
