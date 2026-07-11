package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type MetricsCreateDashboardResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	MetricDashboardSchema *components.MetricDashboardSchema
}

func (m *MetricsCreateDashboardResponse) GetHTTPMeta() components.HTTPMetadata {
	if m == nil {
		return components.HTTPMetadata{}
	}
	return m.HTTPMeta
}

func (m *MetricsCreateDashboardResponse) GetMetricDashboardSchema() *components.MetricDashboardSchema {
	if m == nil {
		return nil
	}
	return m.MetricDashboardSchema
}
