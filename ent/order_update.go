// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent/order"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent/predicate"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent/schema"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetItemTypeExample1 sets the "item_type_example_1" field.
func (ou *OrderUpdate) SetItemTypeExample1(st schema.ItemType) *OrderUpdate {
	ou.mutation.SetItemTypeExample1(st)
	return ou
}

// SetItemTypeExample2 sets the "item_type_example_2" field.
func (ou *OrderUpdate) SetItemTypeExample2(st schema.ItemType) *OrderUpdate {
	ou.mutation.SetItemTypeExample2(st)
	return ou
}

// SetItemTypeExample3 sets the "item_type_example_3" field.
func (ou *OrderUpdate) SetItemTypeExample3(sit schema.StructItemType) *OrderUpdate {
	ou.mutation.SetItemTypeExample3(sit)
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OrderMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.ItemTypeExample1(); ok {
		if err := order.ItemTypeExample1Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_1", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_1": %w`, err)}
		}
	}
	if v, ok := ou.mutation.ItemTypeExample2(); ok {
		if err := order.ItemTypeExample2Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_2", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_2": %w`, err)}
		}
	}
	if v, ok := ou.mutation.ItemTypeExample3(); ok {
		if err := order.ItemTypeExample3Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_3", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_3": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.ItemTypeExample1(); ok {
		_spec.SetField(order.FieldItemTypeExample1, field.TypeEnum, value)
	}
	if value, ok := ou.mutation.ItemTypeExample2(); ok {
		_spec.SetField(order.FieldItemTypeExample2, field.TypeEnum, value)
	}
	if value, ok := ou.mutation.ItemTypeExample3(); ok {
		_spec.SetField(order.FieldItemTypeExample3, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetItemTypeExample1 sets the "item_type_example_1" field.
func (ouo *OrderUpdateOne) SetItemTypeExample1(st schema.ItemType) *OrderUpdateOne {
	ouo.mutation.SetItemTypeExample1(st)
	return ouo
}

// SetItemTypeExample2 sets the "item_type_example_2" field.
func (ouo *OrderUpdateOne) SetItemTypeExample2(st schema.ItemType) *OrderUpdateOne {
	ouo.mutation.SetItemTypeExample2(st)
	return ouo
}

// SetItemTypeExample3 sets the "item_type_example_3" field.
func (ouo *OrderUpdateOne) SetItemTypeExample3(sit schema.StructItemType) *OrderUpdateOne {
	ouo.mutation.SetItemTypeExample3(sit)
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	return withHooks[*Order, OrderMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.ItemTypeExample1(); ok {
		if err := order.ItemTypeExample1Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_1", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_1": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.ItemTypeExample2(); ok {
		if err := order.ItemTypeExample2Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_2", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_2": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.ItemTypeExample3(); ok {
		if err := order.ItemTypeExample3Validator(v); err != nil {
			return &ValidationError{Name: "item_type_example_3", err: fmt.Errorf(`ent: validator failed for field "Order.item_type_example_3": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.ItemTypeExample1(); ok {
		_spec.SetField(order.FieldItemTypeExample1, field.TypeEnum, value)
	}
	if value, ok := ouo.mutation.ItemTypeExample2(); ok {
		_spec.SetField(order.FieldItemTypeExample2, field.TypeEnum, value)
	}
	if value, ok := ouo.mutation.ItemTypeExample3(); ok {
		_spec.SetField(order.FieldItemTypeExample3, field.TypeEnum, value)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}