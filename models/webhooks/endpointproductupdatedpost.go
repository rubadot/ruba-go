package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointproductUpdatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointproductUpdatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointproductUpdatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
