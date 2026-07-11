package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointsubscriptionCreatedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointsubscriptionCreatedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointsubscriptionCreatedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
