// Code generated by ent, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"github.com/go-faster/jx"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent/order"
	"github.com/notmyalias/ogent-custom-ent-enum-issue/ent/schema"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { return &OgentHandler{c} }

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

// CreateOrder handles POST /orders requests.
func (h *OgentHandler) CreateOrder(ctx context.Context, req *CreateOrderReq) (CreateOrderRes, error) {
	b := h.client.Order.Create()
	// Add all fields.
	b.SetItemTypeExample1(schema.ItemType(req.ItemTypeExample1))
	b.SetItemTypeExample2(schema.ItemType(req.ItemTypeExample2))
	b.SetItemTypeExample3(schema.StructItemType(req.ItemTypeExample3))
	// Add all edges.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Order.Query().Where(order.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewOrderCreate(e), nil
}

// ReadOrder handles GET /orders/{id} requests.
func (h *OgentHandler) ReadOrder(ctx context.Context, params ReadOrderParams) (ReadOrderRes, error) {
	q := h.client.Order.Query().Where(order.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewOrderRead(e), nil
}

// UpdateOrder handles PATCH /orders/{id} requests.
func (h *OgentHandler) UpdateOrder(ctx context.Context, req *UpdateOrderReq, params UpdateOrderParams) (UpdateOrderRes, error) {
	b := h.client.Order.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.ItemTypeExample1.Get(); ok {
		b.SetItemTypeExample1(schema.ItemType(v))
	}
	if v, ok := req.ItemTypeExample2.Get(); ok {
		b.SetItemTypeExample2(schema.ItemType(v))
	}
	if v, ok := req.ItemTypeExample3.Get(); ok {
		b.SetItemTypeExample3(schema.StructItemType(v))
	}
	// Add all edges.
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Order.Query().Where(order.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewOrderUpdate(e), nil
}

// DeleteOrder handles DELETE /orders/{id} requests.
func (h *OgentHandler) DeleteOrder(ctx context.Context, params DeleteOrderParams) (DeleteOrderRes, error) {
	err := h.client.Order.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteOrderNoContent), nil

}

// ListOrder handles GET /orders requests.
func (h *OgentHandler) ListOrder(ctx context.Context, params ListOrderParams) (ListOrderRes, error) {
	q := h.client.Order.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewOrderLists(es)
	return (*ListOrderOKApplicationJSON)(&r), nil
}
