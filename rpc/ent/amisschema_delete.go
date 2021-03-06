// Code generated by entc, DO NOT EDIT.

package ent

import (
	"antd-pro-amis-server/rpc/ent/amisschema"
	"antd-pro-amis-server/rpc/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AmisSchemaDelete is the builder for deleting a AmisSchema entity.
type AmisSchemaDelete struct {
	config
	hooks    []Hook
	mutation *AmisSchemaMutation
}

// Where appends a list predicates to the AmisSchemaDelete builder.
func (asd *AmisSchemaDelete) Where(ps ...predicate.AmisSchema) *AmisSchemaDelete {
	asd.mutation.Where(ps...)
	return asd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asd *AmisSchemaDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(asd.hooks) == 0 {
		affected, err = asd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AmisSchemaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asd.mutation = mutation
			affected, err = asd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(asd.hooks) - 1; i >= 0; i-- {
			if asd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (asd *AmisSchemaDelete) ExecX(ctx context.Context) int {
	n, err := asd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asd *AmisSchemaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: amisschema.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: amisschema.FieldID,
			},
		},
	}
	if ps := asd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, asd.driver, _spec)
}

// AmisSchemaDeleteOne is the builder for deleting a single AmisSchema entity.
type AmisSchemaDeleteOne struct {
	asd *AmisSchemaDelete
}

// Exec executes the deletion query.
func (asdo *AmisSchemaDeleteOne) Exec(ctx context.Context) error {
	n, err := asdo.asd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{amisschema.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asdo *AmisSchemaDeleteOne) ExecX(ctx context.Context) {
	asdo.asd.ExecX(ctx)
}
