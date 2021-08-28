package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

// Annotations of the Agent.
func (Agent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "agents"},
	}
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),
		field.Bool("online"),
		field.String("version").MaxLen(32),

		field.String("ipv4").MaxLen(16),
		field.String("ipv6").MaxLen(40),
		field.String("public_ipv4").MaxLen(16),
		field.String("public_ipv6").MaxLen(40),

		field.Time("last_contact").Default(time.Now),
		field.Bool("disabled"),

		field.String("uuid").MaxLen(128).Unique(),
		field.String("machineId").MaxLen(128).Unique(),
		field.String("description"),

		field.Uint("clusterId"),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("node", Node.Type),
		edge.From("owner", Cluster.Type).Ref("agents").Unique(),
		//edge.From("owner", Node.Type).Ref("node").Unique(),
	}
}

// Mixin Agent
func (Agent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Agent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid", "machineId", "clusterId"),
	}
}
