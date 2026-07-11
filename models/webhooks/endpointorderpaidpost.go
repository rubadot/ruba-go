package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointorderPaidPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointorderPaidPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointorderPaidPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
