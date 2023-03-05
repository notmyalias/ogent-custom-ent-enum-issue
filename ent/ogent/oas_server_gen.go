// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateOrder implements createOrder operation.
	//
	// Creates a new Order and persists it to storage.
	//
	// POST /orders
	CreateOrder(ctx context.Context, req *CreateOrderReq) (CreateOrderRes, error)
	// DeleteOrder implements deleteOrder operation.
	//
	// Deletes the Order with the requested ID.
	//
	// DELETE /orders/{id}
	DeleteOrder(ctx context.Context, params DeleteOrderParams) (DeleteOrderRes, error)
	// ListOrder implements listOrder operation.
	//
	// List Orders.
	//
	// GET /orders
	ListOrder(ctx context.Context, params ListOrderParams) (ListOrderRes, error)
	// ReadOrder implements readOrder operation.
	//
	// Finds the Order with the requested ID and returns it.
	//
	// GET /orders/{id}
	ReadOrder(ctx context.Context, params ReadOrderParams) (ReadOrderRes, error)
	// UpdateOrder implements updateOrder operation.
	//
	// Updates a Order and persists changes to storage.
	//
	// PATCH /orders/{id}
	UpdateOrder(ctx context.Context, req *UpdateOrderReq, params UpdateOrderParams) (UpdateOrderRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}