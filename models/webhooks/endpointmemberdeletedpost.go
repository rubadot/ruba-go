package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointmemberDeletedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointmemberDeletedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointmemberDeletedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
