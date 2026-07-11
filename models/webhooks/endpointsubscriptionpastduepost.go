package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointsubscriptionPastDuePostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointsubscriptionPastDuePostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointsubscriptionPastDuePostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
