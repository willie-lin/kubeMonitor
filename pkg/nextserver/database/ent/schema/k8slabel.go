package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// K8sLabel holds the schema definition for the K8sLabel entity.
type K8sLabel struct {
	ent.Schema
}


// Annotations of the Cluster.
func (K8sLabel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "k8sLabels"},
	}
}

// Mixin Container
func (K8sLabel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}



// Fields of the K8sLabel.
func (K8sLabel) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").Unique(),

		field.String("label").MaxLen(256),

		field.Uint("k8sObjectId"),
	}
}

// Edges of the K8sLabel.
func (K8sLabel) Edges() []ent.Edge {
	return nil
}

func (K8sLabel) Indexes() []ent.Index  {
	return []ent.Index {
		index.Fields("label"),
	}

}
