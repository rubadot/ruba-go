package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointcustomerDeletedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointcustomerDeletedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointcustomerDeletedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
