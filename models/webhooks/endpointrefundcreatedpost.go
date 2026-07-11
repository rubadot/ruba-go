package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointrefundCreatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointrefundCreatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointrefundCreatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
