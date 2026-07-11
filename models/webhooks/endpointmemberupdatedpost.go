package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointmemberUpdatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointmemberUpdatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointmemberUpdatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
