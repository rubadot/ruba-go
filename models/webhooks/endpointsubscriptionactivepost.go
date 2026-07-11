package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointsubscriptionActivePostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointsubscriptionActivePostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointsubscriptionActivePostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
