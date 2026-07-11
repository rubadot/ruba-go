package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type OrganizationAccessTokensDeleteRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *OrganizationAccessTokensDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type OrganizationAccessTokensDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *OrganizationAccessTokensDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
