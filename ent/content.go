// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ent-join/ent/content"
	"github.com/ent-join/ent/contentmoviemetadata"
	"github.com/ent-join/ent/uploadedcontent"
)

// Content is the model entity for the Content schema.
type Content struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"filename,omitempty"`
	// UploadedContentFilename holds the value of the "uploaded_content_filename" field.
	UploadedContentFilename string `json:"uploaded_content_filename,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContentQuery when eager-loading is set.
	Edges        ContentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ContentEdges holds the relations/edges for other nodes in the graph.
type ContentEdges struct {
	// UploadedContent holds the value of the uploaded_content edge.
	UploadedContent *UploadedContent `json:"uploaded_content,omitempty"`
	// ContentMovieMetadata holds the value of the content_movie_metadata edge.
	ContentMovieMetadata *ContentMovieMetadata `json:"content_movie_metadata,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UploadedContentOrErr returns the UploadedContent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContentEdges) UploadedContentOrErr() (*UploadedContent, error) {
	if e.loadedTypes[0] {
		if e.UploadedContent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: uploadedcontent.Label}
		}
		return e.UploadedContent, nil
	}
	return nil, &NotLoadedError{edge: "uploaded_content"}
}

// ContentMovieMetadataOrErr returns the ContentMovieMetadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContentEdges) ContentMovieMetadataOrErr() (*ContentMovieMetadata, error) {
	if e.loadedTypes[1] {
		if e.ContentMovieMetadata == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: contentmoviemetadata.Label}
		}
		return e.ContentMovieMetadata, nil
	}
	return nil, &NotLoadedError{edge: "content_movie_metadata"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Content) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case content.FieldID, content.FieldUploadedContentFilename:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Content fields.
func (c *Content) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case content.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case content.FieldUploadedContentFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uploaded_content_filename", values[i])
			} else if value.Valid {
				c.UploadedContentFilename = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Content.
// This includes values selected through modifiers, order, etc.
func (c *Content) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryUploadedContent queries the "uploaded_content" edge of the Content entity.
func (c *Content) QueryUploadedContent() *UploadedContentQuery {
	return NewContentClient(c.config).QueryUploadedContent(c)
}

// QueryContentMovieMetadata queries the "content_movie_metadata" edge of the Content entity.
func (c *Content) QueryContentMovieMetadata() *ContentMovieMetadataQuery {
	return NewContentClient(c.config).QueryContentMovieMetadata(c)
}

// Update returns a builder for updating this Content.
// Note that you need to call Content.Unwrap() before calling this method if this Content
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Content) Update() *ContentUpdateOne {
	return NewContentClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Content entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Content) Unwrap() *Content {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Content is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Content) String() string {
	var builder strings.Builder
	builder.WriteString("Content(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("uploaded_content_filename=")
	builder.WriteString(c.UploadedContentFilename)
	builder.WriteByte(')')
	return builder.String()
}

// Contents is a parsable slice of Content.
type Contents []*Content
