package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type MetersUpdateRequest struct {
	// The meter ID.
	ID          string                 `pathParam:"style=simple,explode=false,name=id"`
	MeterUpdate components.MeterUpdate `request:"mediaType=application/json"`
}

func (m *MetersUpdateRequest) GetID() string {
	if m == nil {
		return ""
	}
	return m.ID
}

func (m *MetersUpdateRequest) GetMeterUpdate() components.MeterUpdate {
	if m == nil {
		return components.MeterUpdate{}
	}
	return m.MeterUpdate
}

type MetersUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Meter updated.
	Meter *components.Meter
}

func (m *MetersUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if m == nil {
		return components.HTTPMetadata{}
	}
	return m.HTTPMeta
}

func (m *MetersUpdateResponse) GetMeter() *components.Meter {
	if m == nil {
		return nil
	}
	return m.Meter
}
