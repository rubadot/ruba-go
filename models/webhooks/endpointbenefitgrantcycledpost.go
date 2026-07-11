package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointbenefitGrantCycledPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointbenefitGrantCycledPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointbenefitGrantCycledPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
