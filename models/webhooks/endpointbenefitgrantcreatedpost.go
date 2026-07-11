package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointbenefitGrantCreatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointbenefitGrantCreatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointbenefitGrantCreatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
