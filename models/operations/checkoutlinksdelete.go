package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CheckoutLinksDeleteRequest struct {
	// The checkout link ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (c *CheckoutLinksDeleteRequest) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

type CheckoutLinksDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (c *CheckoutLinksDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}
