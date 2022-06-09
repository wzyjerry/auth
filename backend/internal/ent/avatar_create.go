// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wzyjerry/auth/internal/ent/avatar"
)

// AvatarCreate is the builder for creating a Avatar entity.
type AvatarCreate struct {
	config
	mutation *AvatarMutation
	hooks    []Hook
}

// SetAvatar sets the "avatar" field.
func (ac *AvatarCreate) SetAvatar(s string) *AvatarCreate {
	ac.mutation.SetAvatar(s)
	return ac
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (ac *AvatarCreate) SetNillableAvatar(s *string) *AvatarCreate {
	if s != nil {
		ac.SetAvatar(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AvatarCreate) SetID(s string) *AvatarCreate {
	ac.mutation.SetID(s)
	return ac
}

// Mutation returns the AvatarMutation object of the builder.
func (ac *AvatarCreate) Mutation() *AvatarMutation {
	return ac.mutation
}

// Save creates the Avatar in the database.
func (ac *AvatarCreate) Save(ctx context.Context) (*Avatar, error) {
	var (
		err  error
		node *Avatar
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AvatarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AvatarCreate) SaveX(ctx context.Context) *Avatar {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AvatarCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AvatarCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AvatarCreate) check() error {
	return nil
}

func (ac *AvatarCreate) sqlSave(ctx context.Context) (*Avatar, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Avatar.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (ac *AvatarCreate) createSpec() (*Avatar, *sqlgraph.CreateSpec) {
	var (
		_node = &Avatar{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: avatar.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: avatar.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: avatar.FieldAvatar,
		})
		_node.Avatar = &value
	}
	return _node, _spec
}

// AvatarCreateBulk is the builder for creating many Avatar entities in bulk.
type AvatarCreateBulk struct {
	config
	builders []*AvatarCreate
}

// Save creates the Avatar entities in the database.
func (acb *AvatarCreateBulk) Save(ctx context.Context) ([]*Avatar, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Avatar, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AvatarMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AvatarCreateBulk) SaveX(ctx context.Context) []*Avatar {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AvatarCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AvatarCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
