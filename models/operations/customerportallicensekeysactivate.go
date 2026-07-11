package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalLicenseKeysActivateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	LicenseKeyActivationRead *components.LicenseKeyActivationRead
}

func (c *CustomerPortalLicenseKeysActivateResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalLicenseKeysActivateResponse) GetLicenseKeyActivationRead() *components.LicenseKeyActivationRead {
	if c == nil {
		return nil
	}
	return c.LicenseKeyActivationRead
}
