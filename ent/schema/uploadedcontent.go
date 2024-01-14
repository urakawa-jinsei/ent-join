package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UploadedContent holds the schema definition for the UploadedContent entity.
type UploadedContent struct {
	ent.Schema
}

func (UploadedContent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "uploaded_content"},
	}
}

// Fields of the UploadedContent.
func (UploadedContent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable().
			StructTag(`json:"filename,omitempty"`).
			StorageKey("filename"),
	}
}

// Edges of the UploadedContent.
func (UploadedContent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contents", Content.Type),
	}
}
