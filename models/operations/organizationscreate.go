package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type OrganizationsCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Organization created.
	Organization *components.Organization
}

func (o *OrganizationsCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *OrganizationsCreateResponse) GetOrganization() *components.Organization {
	if o == nil {
		return nil
	}
	return o.Organization
}
