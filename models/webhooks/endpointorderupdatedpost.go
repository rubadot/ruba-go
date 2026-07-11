package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointorderUpdatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointorderUpdatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointorderUpdatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
