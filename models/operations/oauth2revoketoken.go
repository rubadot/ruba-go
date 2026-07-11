package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type Oauth2RevokeTokenResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	RevokeTokenResponse *components.RevokeTokenResponse
}

func (o *Oauth2RevokeTokenResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *Oauth2RevokeTokenResponse) GetRevokeTokenResponse() *components.RevokeTokenResponse {
	if o == nil {
		return nil
	}
	return o.RevokeTokenResponse
}

// #region class-body-oauth2revoketokenresponse
// #endregion class-body-oauth2revoketokenresponse
