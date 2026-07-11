package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointrefundUpdatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointrefundUpdatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointrefundUpdatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
