package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type LicenseKeysValidateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ValidatedLicenseKey *components.ValidatedLicenseKey
}

func (l *LicenseKeysValidateResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}

func (l *LicenseKeysValidateResponse) GetValidatedLicenseKey() *components.ValidatedLicenseKey {
	if l == nil {
		return nil
	}
	return l.ValidatedLicenseKey
}
