// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateOrder implements createOrder operation.
//
// Creates a new Order and persists it to storage.
//
// POST /orders
func (UnimplementedHandler) CreateOrder(ctx context.Context, req *CreateOrderReq) (r CreateOrderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteOrder implements deleteOrder operation.
//
// Deletes the Order with the requested ID.
//
// DELETE /orders/{id}
func (UnimplementedHandler) DeleteOrder(ctx context.Context, params DeleteOrderParams) (r DeleteOrderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListOrder implements listOrder operation.
//
// List Orders.
//
// GET /orders
func (UnimplementedHandler) ListOrder(ctx context.Context, params ListOrderParams) (r ListOrderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ReadOrder implements readOrder operation.
//
// Finds the Order with the requested ID and returns it.
//
// GET /orders/{id}
func (UnimplementedHandler) ReadOrder(ctx context.Context, params ReadOrderParams) (r ReadOrderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateOrder implements updateOrder operation.
//
// Updates a Order and persists changes to storage.
//
// PATCH /orders/{id}
func (UnimplementedHandler) UpdateOrder(ctx context.Context, req *UpdateOrderReq, params UpdateOrderParams) (r UpdateOrderRes, _ error) {
	return r, ht.ErrNotImplemented
}
