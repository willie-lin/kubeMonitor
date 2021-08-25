package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}


// Annotations of the Cluster.
func (Setting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "settings"},
	}
}

// Mixin Container
func (Setting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}


// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field {
		field.Uint("id").Unique(),

		field.String("name").MaxLen(128),
		field.String("value"),

	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}

func (Setting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}