package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type DiscountsGetRequest struct {
	// The discount ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (d *DiscountsGetRequest) GetID() string {
	if d == nil {
		return ""
	}
	return d.ID
}

type DiscountsGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Discount *components.Discount
}

func (d *DiscountsGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if d == nil {
		return components.HTTPMetadata{}
	}
	return d.HTTPMeta
}

func (d *DiscountsGetResponse) GetDiscount() *components.Discount {
	if d == nil {
		return nil
	}
	return d.Discount
}
