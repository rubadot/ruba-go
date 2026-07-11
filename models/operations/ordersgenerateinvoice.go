package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type OrdersGenerateInvoiceRequest struct {
	// The order ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *OrdersGenerateInvoiceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type OrdersGenerateInvoiceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (o *OrdersGenerateInvoiceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OrdersGenerateInvoiceResponse) GetAny() any {
	if o == nil {
		return nil
	}
	return o.Any
}
