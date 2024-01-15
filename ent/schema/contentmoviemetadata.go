package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ContentMovieMetadata holds the schema definition for the ContentMovieMetadata entity.
type ContentMovieMetadata struct {
	ent.Schema
}

func (ContentMovieMetadata) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "content_movie_metadata"},
	}
}

// Fields of the ContentMovieMetadata.
func (ContentMovieMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable().
			StructTag(`json:"filename,omitempty"`).
			StorageKey("filename"),
		field.Int("width"),
		field.Int("height"),
	}
}

// Edges of the ContentMovieMetadata.
func (ContentMovieMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("content", Content.Type).
			Ref("content_movie_metadata").
			Required().
			Unique(),
	}
}
