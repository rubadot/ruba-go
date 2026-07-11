package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointcustomerSeatAssignedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointcustomerSeatAssignedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointcustomerSeatAssignedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
