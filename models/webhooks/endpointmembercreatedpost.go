package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointmemberCreatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointmemberCreatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointmemberCreatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
