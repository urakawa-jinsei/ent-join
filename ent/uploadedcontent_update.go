// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ent-join/ent/content"
	"github.com/ent-join/ent/predicate"
	"github.com/ent-join/ent/uploadedcontent"
)

// UploadedContentUpdate is the builder for updating UploadedContent entities.
type UploadedContentUpdate struct {
	config
	hooks    []Hook
	mutation *UploadedContentMutation
}

// Where appends a list predicates to the UploadedContentUpdate builder.
func (ucu *UploadedContentUpdate) Where(ps ...predicate.UploadedContent) *UploadedContentUpdate {
	ucu.mutation.Where(ps...)
	return ucu
}

// AddContentIDs adds the "contents" edge to the Content entity by IDs.
func (ucu *UploadedContentUpdate) AddContentIDs(ids ...string) *UploadedContentUpdate {
	ucu.mutation.AddContentIDs(ids...)
	return ucu
}

// AddContents adds the "contents" edges to the Content entity.
func (ucu *UploadedContentUpdate) AddContents(c ...*Content) *UploadedContentUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ucu.AddContentIDs(ids...)
}

// Mutation returns the UploadedContentMutation object of the builder.
func (ucu *UploadedContentUpdate) Mutation() *UploadedContentMutation {
	return ucu.mutation
}

// ClearContents clears all "contents" edges to the Content entity.
func (ucu *UploadedContentUpdate) ClearContents() *UploadedContentUpdate {
	ucu.mutation.ClearContents()
	return ucu
}

// RemoveContentIDs removes the "contents" edge to Content entities by IDs.
func (ucu *UploadedContentUpdate) RemoveContentIDs(ids ...string) *UploadedContentUpdate {
	ucu.mutation.RemoveContentIDs(ids...)
	return ucu
}

// RemoveContents removes "contents" edges to Content entities.
func (ucu *UploadedContentUpdate) RemoveContents(c ...*Content) *UploadedContentUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ucu.RemoveContentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucu *UploadedContentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ucu.sqlSave, ucu.mutation, ucu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucu *UploadedContentUpdate) SaveX(ctx context.Context) int {
	affected, err := ucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucu *UploadedContentUpdate) Exec(ctx context.Context) error {
	_, err := ucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucu *UploadedContentUpdate) ExecX(ctx context.Context) {
	if err := ucu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ucu *UploadedContentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(uploadedcontent.Table, uploadedcontent.Columns, sqlgraph.NewFieldSpec(uploadedcontent.FieldID, field.TypeString))
	if ps := ucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ucu.mutation.ContentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   uploadedcontent.ContentsTable,
			Columns: []string{uploadedcontent.ContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(content.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucu.mutation.RemovedContentsIDs(); len(nodes) > 0 && !ucu.mutation.ContentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   uploadedcontent.ContentsTable,
			Columns: []string{uploadedcontent.ContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(content.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucu.mutation.ContentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   uploadedcontent.ContentsTable,
			Columns: []string{uploadedcontent.ContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(content.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uploadedcontent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ucu.mutation.done = true
	return n, nil
}

// UploadedContentUpdateOne is the builder for updating a single UploadedContent entity.
type UploadedContentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UploadedContentMutation
}

// AddContentIDs adds the "contents" edge to the Content entity by IDs.
func (ucuo *UploadedContentUpdateOne) AddContentIDs(ids ...string) *UploadedContentUpdateOne {
	ucuo.mutation.AddContentIDs(ids...)
	return ucuo
}

// AddContents adds the "contents" edges to the Content entity.
func (ucuo *UploadedContentUpdateOne) AddContents(c ...*Content) *UploadedContentUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ucuo.AddContentIDs(ids...)
}

// Mutation returns the UploadedContentMutation object of the builder.
func (ucuo *UploadedContentUpdateOne) Mutation() *UploadedContentMutation {
	return ucuo.mutation
}

// ClearContents clears all "contents" edges to the Content entity.
func (ucuo *UploadedContentUpdateOne) ClearContents() *UploadedContentUpdateOne {
	ucuo.mutation.ClearContents()
	return ucuo
}

// RemoveContentIDs removes the "contents" edge to Content entities by IDs.
func (ucuo *UploadedContentUpdateOne) RemoveContentIDs(ids ...string) *UploadedContentUpdateOne {
	ucuo.mutation.RemoveContentIDs(ids...)
	return ucuo
}

// RemoveContents removes "contents" edges to Content entities.
func (ucuo *UploadedContentUpdateOne) RemoveContents(c ...*Content) *UploadedContentUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ucuo.RemoveContentIDs(ids...)
}

// Where appends a list predicates to the UploadedContentUpdate builder.
func (ucuo *UploadedContentUpdateOne) Where(ps ...predicate.UploadedContent) *UploadedContentUpdateOne {
	ucuo.mutation.Where(ps...)
	return ucuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucuo *UploadedContentUpdateOne) Select(field string, fields ...string) *UploadedContentUpdateOne {
	ucuo.fields = append([]string{field}, fields...)
	return ucuo
}

// Save executes the query and returns the updated UploadedContent entity.
func (ucuo *UploadedContentUpdateOne) Save(ctx context.Context) (*UploadedContent, error) {
	return withHooks(ctx, ucuo.sqlSave, ucuo.mutation, ucuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucuo *UploadedContentUpdateOne) SaveX(ctx context.Context) *UploadedContent {
	node, err := ucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucuo *UploadedContentUpdateOne) Exec(ctx context.Context) error {
	_, err := ucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucuo *UploadedContentUpdateOne) ExecX(ctx context.Context) {
	if err := ucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ucuo *UploadedContentUpdateOne) sqlSave(ctx context.Context) (_node *UploadedContent, err error) {
	_spec := sqlgraph.NewUpdateSpec(uploadedcontent.Table, uploadedcontent.Columns, sqlgraph.NewFieldSpec(uploadedcontent.FieldID, field.TypeString))
	id, ok := ucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UploadedContent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uploadedcontent.FieldID)
		for _, f := range fields {
			if !uploadedcontent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != uploadedcontent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ucuo.mutation.ContentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   uploadedcontent.ContentsTable,
			Columns: []string{uploadedcontent.ContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(content.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucuo.mutation.RemovedContentsIDs(); len(nodes) > 0 && !ucuo.mutation.ContentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   uploadedcontent.ContentsTable,
			Columns: []string{uploadedcontent.ContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(content.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucuo.mutation.ContentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   uploadedcontent.ContentsTable,
			Columns: []string{uploadedcontent.ContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(content.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UploadedContent{config: ucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uploadedcontent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucuo.mutation.done = true
	return _node, nil
}
