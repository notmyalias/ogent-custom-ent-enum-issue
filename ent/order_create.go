// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent/order"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent/schema"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetItemTypeExample1 sets the "item_type_example_1" field.
func (oc *OrderCreate) SetItemTypeExample1(st schema.ItemType) *OrderCreate {
	oc.mutation.SetItemTypeExample1(st)
	return oc
}

// SetItemTypeExample2 sets the "item_type_example_2" field.
func (oc *OrderCreate) SetItemTypeExample2(st schema.ItemType) *OrderCreate {
	oc.mutation.SetItemTypeExample2(st)
	return oc
}

// SetItemTypeExample3 sets the "item_type_example_3" field.
func (oc *OrderCreate) SetItemTypeExample3(sit schema.StructItemType) *OrderCreate {
	oc.mutation.SetItemTypeExample3(sit)
	return oc
}

// SetID sets the "id" field.
func (oc *OrderCreate) SetID(u uuid.UUID) *OrderCreate {
	oc.mutation.SetID(u)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableID(u *uuid.UUID) *OrderCreate {
	if u != nil {
		oc.SetID(*u)
	}
	return oc
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks[*Order, OrderMutation](ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.ID(); !ok {
		v := order.DefaultID()
		oc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.ItemTypeExample1(); !ok {
		return &ValidationError{Name: "item_type_example_1", err: errors.New(`ent: missing required field "Order.item_type_example_1"`)}
	}
	if v, ok := oc.mutation.ItemTypeExample1(); ok {
		if err := order.ItemTypeExample1Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_1", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_1": %w`, err)}
		}
	}
	if _, ok := oc.mutation.ItemTypeExample2(); !ok {
		return &ValidationError{Name: "item_type_example_2", err: errors.New(`ent: missing required field "Order.item_type_example_2"`)}
	}
	if v, ok := oc.mutation.ItemTypeExample2(); ok {
		if err := order.ItemTypeExample2Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_2", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_2": %w`, err)}
		}
	}
	if _, ok := oc.mutation.ItemTypeExample3(); !ok {
		return &ValidationError{Name: "item_type_example_3", err: errors.New(`ent: missing required field "Order.item_type_example_3"`)}
	}
	if v, ok := oc.mutation.ItemTypeExample3(); ok {
		if err := order.ItemTypeExample3Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_3", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_3": %w`, err)}
		}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oc.mutation.ItemTypeExample1(); ok {
		_spec.SetField(order.FieldItemTypeExample1, field.TypeEnum, value)
		_node.ItemTypeExample1 = value
	}
	if value, ok := oc.mutation.ItemTypeExample2(); ok {
		_spec.SetField(order.FieldItemTypeExample2, field.TypeEnum, value)
		_node.ItemTypeExample2 = value
	}
	if value, ok := oc.mutation.ItemTypeExample3(); ok {
		_spec.SetField(order.FieldItemTypeExample3, field.TypeEnum, value)
		_node.ItemTypeExample3 = value
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
