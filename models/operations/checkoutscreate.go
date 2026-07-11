package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CheckoutsCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Checkout session created.
	Checkout *components.Checkout
}

func (c *CheckoutsCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CheckoutsCreateResponse) GetCheckout() *components.Checkout {
	if c == nil {
		return nil
	}
	return c.Checkout
}
