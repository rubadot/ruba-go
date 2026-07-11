package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type OrganizationAccessTokensCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	OrganizationAccessTokenCreateResponse *components.OrganizationAccessTokenCreateResponse
}

func (o *OrganizationAccessTokensCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OrganizationAccessTokensCreateResponse) GetOrganizationAccessTokenCreateResponse() *components.OrganizationAccessTokenCreateResponse {
	if o == nil {
		return nil
	}
	return o.OrganizationAccessTokenCreateResponse
}
