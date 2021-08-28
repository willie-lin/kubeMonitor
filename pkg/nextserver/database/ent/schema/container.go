package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Container holds the schema definition for the Container entity.
type Container struct {
	ent.Schema
}

// Annotations of the Container.
func (Container) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "containers"},
	}
}

// Fields of the Container.
func (Container) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),
		field.String("type"),
		field.String("containerId").MaxLen(128),
		field.String("name").MaxLen(256),
		field.String("image").MaxLen(128),
		field.JSON("info", []string{}),
		field.Uint("clusterId"),
		field.Uint("nodeId"),
	}
}

// Edges of the Container.
func (Container) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("process", Proces.Type),
		edge.From("owner", Node.Type).Ref("containers").Unique(),
	}
}

// Mixin Container
func (Container) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Container) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("containerId", "clusterId", "nodeId"),
	}
}
