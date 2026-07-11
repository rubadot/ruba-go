package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointsubscriptionUncanceledPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointsubscriptionUncanceledPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointsubscriptionUncanceledPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
