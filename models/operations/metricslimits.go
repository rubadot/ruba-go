package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type MetricsLimitsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	MetricsLimits *components.MetricsLimits
}

func (m *MetricsLimitsResponse) GetHTTPMeta() components.HTTPMetadata {
	if m == nil {
		return components.HTTPMetadata{}
	}
	return m.HTTPMeta
}

func (m *MetricsLimitsResponse) GetMetricsLimits() *components.MetricsLimits {
	if m == nil {
		return nil
	}
	return m.MetricsLimits
}
