// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ent-join/ent/predicate"
	"github.com/ent-join/ent/uploadedcontent"
)

// UploadedContentDelete is the builder for deleting a UploadedContent entity.
type UploadedContentDelete struct {
	config
	hooks    []Hook
	mutation *UploadedContentMutation
}

// Where appends a list predicates to the UploadedContentDelete builder.
func (ucd *UploadedContentDelete) Where(ps ...predicate.UploadedContent) *UploadedContentDelete {
	ucd.mutation.Where(ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UploadedContentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ucd.sqlExec, ucd.mutation, ucd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UploadedContentDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UploadedContentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(uploadedcontent.Table, sqlgraph.NewFieldSpec(uploadedcontent.FieldID, field.TypeString))
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ucd.mutation.done = true
	return affected, err
}

// UploadedContentDeleteOne is the builder for deleting a single UploadedContent entity.
type UploadedContentDeleteOne struct {
	ucd *UploadedContentDelete
}

// Where appends a list predicates to the UploadedContentDelete builder.
func (ucdo *UploadedContentDeleteOne) Where(ps ...predicate.UploadedContent) *UploadedContentDeleteOne {
	ucdo.ucd.mutation.Where(ps...)
	return ucdo
}

// Exec executes the deletion query.
func (ucdo *UploadedContentDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{uploadedcontent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UploadedContentDeleteOne) ExecX(ctx context.Context) {
	if err := ucdo.Exec(ctx); err != nil {
		panic(err)
	}
}
