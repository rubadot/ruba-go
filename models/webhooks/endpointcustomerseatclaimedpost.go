package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointcustomerSeatClaimedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointcustomerSeatClaimedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointcustomerSeatClaimedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
