package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalLicenseKeysDeactivateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (c *CustomerPortalLicenseKeysDeactivateResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}
