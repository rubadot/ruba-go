package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointcustomerCreatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointcustomerCreatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointcustomerCreatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
