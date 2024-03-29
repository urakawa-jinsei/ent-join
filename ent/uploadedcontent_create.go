// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ent-join/ent/content"
	"github.com/ent-join/ent/uploadedcontent"
)

// UploadedContentCreate is the builder for creating a UploadedContent entity.
type UploadedContentCreate struct {
	config
	mutation *UploadedContentMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (ucc *UploadedContentCreate) SetID(s string) *UploadedContentCreate {
	ucc.mutation.SetID(s)
	return ucc
}

// AddContentIDs adds the "contents" edge to the Content entity by IDs.
func (ucc *UploadedContentCreate) AddContentIDs(ids ...string) *UploadedContentCreate {
	ucc.mutation.AddContentIDs(ids...)
	return ucc
}

// AddContents adds the "contents" edges to the Content entity.
func (ucc *UploadedContentCreate) AddContents(c ...*Content) *UploadedContentCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ucc.AddContentIDs(ids...)
}

// Mutation returns the UploadedContentMutation object of the builder.
func (ucc *UploadedContentCreate) Mutation() *UploadedContentMutation {
	return ucc.mutation
}

// Save creates the UploadedContent in the database.
func (ucc *UploadedContentCreate) Save(ctx context.Context) (*UploadedContent, error) {
	return withHooks(ctx, ucc.sqlSave, ucc.mutation, ucc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ucc *UploadedContentCreate) SaveX(ctx context.Context) *UploadedContent {
	v, err := ucc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucc *UploadedContentCreate) Exec(ctx context.Context) error {
	_, err := ucc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucc *UploadedContentCreate) ExecX(ctx context.Context) {
	if err := ucc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucc *UploadedContentCreate) check() error {
	if v, ok := ucc.mutation.ID(); ok {
		if err := uploadedcontent.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "UploadedContent.id": %w`, err)}
		}
	}
	return nil
}

func (ucc *UploadedContentCreate) sqlSave(ctx context.Context) (*UploadedContent, error) {
	if err := ucc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ucc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UploadedContent.ID type: %T", _spec.ID.Value)
		}
	}
	ucc.mutation.id = &_node.ID
	ucc.mutation.done = true
	return _node, nil
}

func (ucc *UploadedContentCreate) createSpec() (*UploadedContent, *sqlgraph.CreateSpec) {
	var (
		_node = &UploadedContent{config: ucc.config}
		_spec = sqlgraph.NewCreateSpec(uploadedcontent.Table, sqlgraph.NewFieldSpec(uploadedcontent.FieldID, field.TypeString))
	)
	if id, ok := ucc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := ucc.mutation.ContentsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UploadedContentCreateBulk is the builder for creating many UploadedContent entities in bulk.
type UploadedContentCreateBulk struct {
	config
	err      error
	builders []*UploadedContentCreate
}

// Save creates the UploadedContent entities in the database.
func (uccb *UploadedContentCreateBulk) Save(ctx context.Context) ([]*UploadedContent, error) {
	if uccb.err != nil {
		return nil, uccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uccb.builders))
	nodes := make([]*UploadedContent, len(uccb.builders))
	mutators := make([]Mutator, len(uccb.builders))
	for i := range uccb.builders {
		func(i int, root context.Context) {
			builder := uccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UploadedContentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uccb *UploadedContentCreateBulk) SaveX(ctx context.Context) []*UploadedContent {
	v, err := uccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uccb *UploadedContentCreateBulk) Exec(ctx context.Context) error {
	_, err := uccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uccb *UploadedContentCreateBulk) ExecX(ctx context.Context) {
	if err := uccb.Exec(ctx); err != nil {
		panic(err)
	}
}
