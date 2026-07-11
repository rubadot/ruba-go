package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type OrdersUpdateRequest struct {
	// The order ID.
	ID          string                 `pathParam:"style=simple,explode=false,name=id"`
	OrderUpdate components.OrderUpdate `request:"mediaType=application/json"`
}

func (o *OrdersUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OrdersUpdateRequest) GetOrderUpdate() components.OrderUpdate {
	if o == nil {
		return components.OrderUpdate{}
	}
	return o.OrderUpdate
}

type OrdersUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Order *components.Order
}

func (o *OrdersUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OrdersUpdateResponse) GetOrder() *components.Order {
	if o == nil {
		return nil
	}
	return o.Order
}
