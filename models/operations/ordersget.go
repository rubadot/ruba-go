package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type OrdersGetRequest struct {
	// The order ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *OrdersGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type OrdersGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Order *components.Order
}

func (o *OrdersGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OrdersGetResponse) GetOrder() *components.Order {
	if o == nil {
		return nil
	}
	return o.Order
}
