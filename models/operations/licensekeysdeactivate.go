package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type LicenseKeysDeactivateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (l *LicenseKeysDeactivateResponse) GetHTTPMeta() components.HTTPMetadata {
	if l == nil {
		return components.HTTPMetadata{}
	}
	return l.HTTPMeta
}
