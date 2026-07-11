package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointsubscriptionCanceledPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointsubscriptionCanceledPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointsubscriptionCanceledPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
