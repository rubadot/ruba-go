package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointcustomerUpdatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointcustomerUpdatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointcustomerUpdatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
