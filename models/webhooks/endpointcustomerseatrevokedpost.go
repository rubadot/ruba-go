package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointcustomerSeatRevokedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointcustomerSeatRevokedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointcustomerSeatRevokedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
