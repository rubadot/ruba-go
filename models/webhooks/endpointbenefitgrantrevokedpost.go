package webhooks

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type EndpointbenefitGrantRevokedPostResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	Any any
}

func (e *EndpointbenefitGrantRevokedPostResponse) GetHTTPMeta() components.HTTPMetadata {
	if e == nil {
		return components.HTTPMetadata{}
	}
	return e.HTTPMeta
}

func (e *EndpointbenefitGrantRevokedPostResponse) GetAny() any {
	if e == nil {
		return nil
	}
	return e.Any
}
