package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type LicenseKeysActivateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	LicenseKeyActivationRead *components.LicenseKeyActivationRead
}

func (l *LicenseKeysActivateResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *LicenseKeysActivateResponse) GetLicenseKeyActivationRead() *components.LicenseKeyActivationRead {
	if l == nil {
		return nil
	}
	return l.LicenseKeyActivationRead
}
