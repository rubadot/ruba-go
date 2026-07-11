package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type Oauth2IntrospectTokenResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	IntrospectTokenResponse *components.IntrospectTokenResponse
}

func (o *Oauth2IntrospectTokenResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *Oauth2IntrospectTokenResponse) GetIntrospectTokenResponse() *components.IntrospectTokenResponse {
	if o == nil {
		return nil
	}
	return o.IntrospectTokenResponse
}

// #region class-body-oauth2introspecttokenresponse
// #endregion class-body-oauth2introspecttokenresponse
