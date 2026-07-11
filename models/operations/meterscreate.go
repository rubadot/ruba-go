package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type MetersCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Meter created.
	Meter *components.Meter
}

func (m *MetersCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if m == nil {
		return components.HTTPMetadata{}
	}
	return m.HTTPMeta
}

func (m *MetersCreateResponse) GetMeter() *components.Meter {
	if m == nil {
		return nil
	}
	return m.Meter
}
