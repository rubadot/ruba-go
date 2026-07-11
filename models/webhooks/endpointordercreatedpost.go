package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointorderCreatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointorderCreatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointorderCreatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
