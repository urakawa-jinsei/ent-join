package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Content holds the schema definition for the Content entity.
type Content struct {
	ent.Schema
}

// Content holds the schema definition for the Content entity.
func (Content) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "content"},
	}
}

// Fields of the Content.
func (Content) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable().
			StructTag(`json:"filename,omitempty"`).
			StorageKey("filename"),
		field.String("uploaded_content_filename"),
	}
}

// Edges of the Content.
func (Content) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("uploaded_content", UploadedContent.Type).
			Field("uploaded_content_filename").
			Ref("contents").
			Unique().
			Required(),
		edge.To("content_movie_metadata", ContentMovieMetadata.Type).
			StorageKey(edge.Column("filename")).
			Unique(),
	}
}
