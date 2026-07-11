package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointproductCreatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointproductCreatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointproductCreatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
