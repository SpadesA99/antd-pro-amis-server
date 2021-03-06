// Code generated by entc, DO NOT EDIT.

package ent

import (
	"antd-pro-amis-server/rpc/ent/amisschema"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AmisSchemaCreate is the builder for creating a AmisSchema entity.
type AmisSchemaCreate struct {
	config
	mutation *AmisSchemaMutation
	hooks    []Hook
}

// SetRouter sets the "router" field.
func (asc *AmisSchemaCreate) SetRouter(s string) *AmisSchemaCreate {
	asc.mutation.SetRouter(s)
	return asc
}

// SetSchema sets the "schema" field.
func (asc *AmisSchemaCreate) SetSchema(s string) *AmisSchemaCreate {
	asc.mutation.SetSchema(s)
	return asc
}

// SetCreatedAt sets the "created_at" field.
func (asc *AmisSchemaCreate) SetCreatedAt(t time.Time) *AmisSchemaCreate {
	asc.mutation.SetCreatedAt(t)
	return asc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asc *AmisSchemaCreate) SetNillableCreatedAt(t *time.Time) *AmisSchemaCreate {
	if t != nil {
		asc.SetCreatedAt(*t)
	}
	return asc
}

// SetUpdatedAt sets the "updated_at" field.
func (asc *AmisSchemaCreate) SetUpdatedAt(t time.Time) *AmisSchemaCreate {
	asc.mutation.SetUpdatedAt(t)
	return asc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (asc *AmisSchemaCreate) SetNillableUpdatedAt(t *time.Time) *AmisSchemaCreate {
	if t != nil {
		asc.SetUpdatedAt(*t)
	}
	return asc
}

// Mutation returns the AmisSchemaMutation object of the builder.
func (asc *AmisSchemaCreate) Mutation() *AmisSchemaMutation {
	return asc.mutation
}

// Save creates the AmisSchema in the database.
func (asc *AmisSchemaCreate) Save(ctx context.Context) (*AmisSchema, error) {
	var (
		err  error
		node *AmisSchema
	)
	asc.defaults()
	if len(asc.hooks) == 0 {
		if err = asc.check(); err != nil {
			return nil, err
		}
		node, err = asc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AmisSchemaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = asc.check(); err != nil {
				return nil, err
			}
			asc.mutation = mutation
			if node, err = asc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(asc.hooks) - 1; i >= 0; i-- {
			if asc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (asc *AmisSchemaCreate) SaveX(ctx context.Context) *AmisSchema {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *AmisSchemaCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *AmisSchemaCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asc *AmisSchemaCreate) defaults() {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		v := amisschema.DefaultCreatedAt()
		asc.mutation.SetCreatedAt(v)
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		v := amisschema.DefaultUpdatedAt()
		asc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asc *AmisSchemaCreate) check() error {
	if _, ok := asc.mutation.Router(); !ok {
		return &ValidationError{Name: "router", err: errors.New(`ent: missing required field "AmisSchema.router"`)}
	}
	if _, ok := asc.mutation.Schema(); !ok {
		return &ValidationError{Name: "schema", err: errors.New(`ent: missing required field "AmisSchema.schema"`)}
	}
	if _, ok := asc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AmisSchema.created_at"`)}
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AmisSchema.updated_at"`)}
	}
	return nil
}

func (asc *AmisSchemaCreate) sqlSave(ctx context.Context) (*AmisSchema, error) {
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (asc *AmisSchemaCreate) createSpec() (*AmisSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &AmisSchema{config: asc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: amisschema.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: amisschema.FieldID,
			},
		}
	)
	if value, ok := asc.mutation.Router(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: amisschema.FieldRouter,
		})
		_node.Router = value
	}
	if value, ok := asc.mutation.Schema(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: amisschema.FieldSchema,
		})
		_node.Schema = value
	}
	if value, ok := asc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: amisschema.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := asc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: amisschema.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// AmisSchemaCreateBulk is the builder for creating many AmisSchema entities in bulk.
type AmisSchemaCreateBulk struct {
	config
	builders []*AmisSchemaCreate
}

// Save creates the AmisSchema entities in the database.
func (ascb *AmisSchemaCreateBulk) Save(ctx context.Context) ([]*AmisSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*AmisSchema, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AmisSchemaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *AmisSchemaCreateBulk) SaveX(ctx context.Context) []*AmisSchema {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *AmisSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *AmisSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}
