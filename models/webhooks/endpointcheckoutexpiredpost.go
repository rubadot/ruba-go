package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointcheckoutExpiredPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointcheckoutExpiredPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointcheckoutExpiredPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
