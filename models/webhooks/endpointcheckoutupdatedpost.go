package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointcheckoutUpdatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointcheckoutUpdatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointcheckoutUpdatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
