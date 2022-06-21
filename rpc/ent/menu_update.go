// Code generated by entc, DO NOT EDIT.

package ent

import (
	"antd-pro-amis-server/rpc/ent/menu"
	"antd-pro-amis-server/rpc/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks    []Hook
	mutation *MenuMutation
}

// Where appends a list predicates to the MenuUpdate builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetMenuName sets the "menu_name" field.
func (mu *MenuUpdate) SetMenuName(s string) *MenuUpdate {
	mu.mutation.SetMenuName(s)
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MenuUpdate) SetParentID(i int64) *MenuUpdate {
	mu.mutation.ResetParentID()
	mu.mutation.SetParentID(i)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableParentID(i *int64) *MenuUpdate {
	if i != nil {
		mu.SetParentID(*i)
	}
	return mu
}

// AddParentID adds i to the "parent_id" field.
func (mu *MenuUpdate) AddParentID(i int64) *MenuUpdate {
	mu.mutation.AddParentID(i)
	return mu
}

// SetPath sets the "path" field.
func (mu *MenuUpdate) SetPath(s string) *MenuUpdate {
	mu.mutation.SetPath(s)
	return mu
}

// SetIcon sets the "icon" field.
func (mu *MenuUpdate) SetIcon(s string) *MenuUpdate {
	mu.mutation.SetIcon(s)
	return mu
}

// SetRoles sets the "roles" field.
func (mu *MenuUpdate) SetRoles(s []string) *MenuUpdate {
	mu.mutation.SetRoles(s)
	return mu
}

// SetLayout sets the "layout" field.
func (mu *MenuUpdate) SetLayout(b bool) *MenuUpdate {
	mu.mutation.SetLayout(b)
	return mu
}

// SetStatus sets the "status" field.
func (mu *MenuUpdate) SetStatus(b bool) *MenuUpdate {
	mu.mutation.SetStatus(b)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableStatus(b *bool) *MenuUpdate {
	if b != nil {
		mu.SetStatus(*b)
	}
	return mu
}

// SetHideInMenu sets the "hide_in_menu" field.
func (mu *MenuUpdate) SetHideInMenu(b bool) *MenuUpdate {
	mu.mutation.SetHideInMenu(b)
	return mu
}

// SetSort sets the "sort" field.
func (mu *MenuUpdate) SetSort(i int64) *MenuUpdate {
	mu.mutation.ResetSort()
	mu.mutation.SetSort(i)
	return mu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableSort(i *int64) *MenuUpdate {
	if i != nil {
		mu.SetSort(*i)
	}
	return mu
}

// AddSort adds i to the "sort" field.
func (mu *MenuUpdate) AddSort(i int64) *MenuUpdate {
	mu.mutation.AddSort(i)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MenuUpdate) SetUpdatedAt(t time.Time) *MenuUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MenuUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MenuUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MenuUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MenuUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := menu.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: menu.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.MenuName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldMenuName,
		})
	}
	if value, ok := mu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldParentID,
		})
	}
	if value, ok := mu.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldParentID,
		})
	}
	if value, ok := mu.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldPath,
		})
	}
	if value, ok := mu.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldIcon,
		})
	}
	if value, ok := mu.mutation.Roles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: menu.FieldRoles,
		})
	}
	if value, ok := mu.mutation.Layout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldLayout,
		})
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldStatus,
		})
	}
	if value, ok := mu.mutation.HideInMenu(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldHideInMenu,
		})
	}
	if value, ok := mu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldSort,
		})
	}
	if value, ok := mu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldSort,
		})
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: menu.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MenuUpdateOne is the builder for updating a single Menu entity.
type MenuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MenuMutation
}

// SetMenuName sets the "menu_name" field.
func (muo *MenuUpdateOne) SetMenuName(s string) *MenuUpdateOne {
	muo.mutation.SetMenuName(s)
	return muo
}

// SetParentID sets the "parent_id" field.
func (muo *MenuUpdateOne) SetParentID(i int64) *MenuUpdateOne {
	muo.mutation.ResetParentID()
	muo.mutation.SetParentID(i)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableParentID(i *int64) *MenuUpdateOne {
	if i != nil {
		muo.SetParentID(*i)
	}
	return muo
}

// AddParentID adds i to the "parent_id" field.
func (muo *MenuUpdateOne) AddParentID(i int64) *MenuUpdateOne {
	muo.mutation.AddParentID(i)
	return muo
}

// SetPath sets the "path" field.
func (muo *MenuUpdateOne) SetPath(s string) *MenuUpdateOne {
	muo.mutation.SetPath(s)
	return muo
}

// SetIcon sets the "icon" field.
func (muo *MenuUpdateOne) SetIcon(s string) *MenuUpdateOne {
	muo.mutation.SetIcon(s)
	return muo
}

// SetRoles sets the "roles" field.
func (muo *MenuUpdateOne) SetRoles(s []string) *MenuUpdateOne {
	muo.mutation.SetRoles(s)
	return muo
}

// SetLayout sets the "layout" field.
func (muo *MenuUpdateOne) SetLayout(b bool) *MenuUpdateOne {
	muo.mutation.SetLayout(b)
	return muo
}

// SetStatus sets the "status" field.
func (muo *MenuUpdateOne) SetStatus(b bool) *MenuUpdateOne {
	muo.mutation.SetStatus(b)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableStatus(b *bool) *MenuUpdateOne {
	if b != nil {
		muo.SetStatus(*b)
	}
	return muo
}

// SetHideInMenu sets the "hide_in_menu" field.
func (muo *MenuUpdateOne) SetHideInMenu(b bool) *MenuUpdateOne {
	muo.mutation.SetHideInMenu(b)
	return muo
}

// SetSort sets the "sort" field.
func (muo *MenuUpdateOne) SetSort(i int64) *MenuUpdateOne {
	muo.mutation.ResetSort()
	muo.mutation.SetSort(i)
	return muo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableSort(i *int64) *MenuUpdateOne {
	if i != nil {
		muo.SetSort(*i)
	}
	return muo
}

// AddSort adds i to the "sort" field.
func (muo *MenuUpdateOne) AddSort(i int64) *MenuUpdateOne {
	muo.mutation.AddSort(i)
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MenuUpdateOne) SetUpdatedAt(t time.Time) *MenuUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MenuUpdateOne) Select(field string, fields ...string) *MenuUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Menu entity.
func (muo *MenuUpdateOne) Save(ctx context.Context) (*Menu, error) {
	var (
		err  error
		node *Menu
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MenuUpdateOne) SaveX(ctx context.Context) *Menu {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MenuUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MenuUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MenuUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := menu.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: menu.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Menu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menu.FieldID)
		for _, f := range fields {
			if !menu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.MenuName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldMenuName,
		})
	}
	if value, ok := muo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldParentID,
		})
	}
	if value, ok := muo.mutation.AddedParentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldParentID,
		})
	}
	if value, ok := muo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldPath,
		})
	}
	if value, ok := muo.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldIcon,
		})
	}
	if value, ok := muo.mutation.Roles(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: menu.FieldRoles,
		})
	}
	if value, ok := muo.mutation.Layout(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldLayout,
		})
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldStatus,
		})
	}
	if value, ok := muo.mutation.HideInMenu(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldHideInMenu,
		})
	}
	if value, ok := muo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldSort,
		})
	}
	if value, ok := muo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: menu.FieldSort,
		})
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: menu.FieldUpdatedAt,
		})
	}
	_node = &Menu{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
