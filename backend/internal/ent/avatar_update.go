// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wzyjerry/auth/internal/ent/avatar"
	"github.com/wzyjerry/auth/internal/ent/predicate"
)

// AvatarUpdate is the builder for updating Avatar entities.
type AvatarUpdate struct {
	config
	hooks    []Hook
	mutation *AvatarMutation
}

// Where appends a list predicates to the AvatarUpdate builder.
func (au *AvatarUpdate) Where(ps ...predicate.Avatar) *AvatarUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetAvatar sets the "avatar" field.
func (au *AvatarUpdate) SetAvatar(s string) *AvatarUpdate {
	au.mutation.SetAvatar(s)
	return au
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (au *AvatarUpdate) SetNillableAvatar(s *string) *AvatarUpdate {
	if s != nil {
		au.SetAvatar(*s)
	}
	return au
}

// ClearAvatar clears the value of the "avatar" field.
func (au *AvatarUpdate) ClearAvatar() *AvatarUpdate {
	au.mutation.ClearAvatar()
	return au
}

// Mutation returns the AvatarMutation object of the builder.
func (au *AvatarUpdate) Mutation() *AvatarMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AvatarUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AvatarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AvatarUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AvatarUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AvatarUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AvatarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   avatar.Table,
			Columns: avatar.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: avatar.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: avatar.FieldAvatar,
		})
	}
	if au.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: avatar.FieldAvatar,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{avatar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AvatarUpdateOne is the builder for updating a single Avatar entity.
type AvatarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AvatarMutation
}

// SetAvatar sets the "avatar" field.
func (auo *AvatarUpdateOne) SetAvatar(s string) *AvatarUpdateOne {
	auo.mutation.SetAvatar(s)
	return auo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillableAvatar(s *string) *AvatarUpdateOne {
	if s != nil {
		auo.SetAvatar(*s)
	}
	return auo
}

// ClearAvatar clears the value of the "avatar" field.
func (auo *AvatarUpdateOne) ClearAvatar() *AvatarUpdateOne {
	auo.mutation.ClearAvatar()
	return auo
}

// Mutation returns the AvatarMutation object of the builder.
func (auo *AvatarUpdateOne) Mutation() *AvatarMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AvatarUpdateOne) Select(field string, fields ...string) *AvatarUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Avatar entity.
func (auo *AvatarUpdateOne) Save(ctx context.Context) (*Avatar, error) {
	var (
		err  error
		node *Avatar
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AvatarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AvatarUpdateOne) SaveX(ctx context.Context) *Avatar {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AvatarUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AvatarUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AvatarUpdateOne) sqlSave(ctx context.Context) (_node *Avatar, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   avatar.Table,
			Columns: avatar.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: avatar.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Avatar.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, avatar.FieldID)
		for _, f := range fields {
			if !avatar.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != avatar.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: avatar.FieldAvatar,
		})
	}
	if auo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: avatar.FieldAvatar,
		})
	}
	_node = &Avatar{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{avatar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
