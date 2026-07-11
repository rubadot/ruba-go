package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalMembersAddMemberResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Member added.
	CustomerPortalMember *components.CustomerPortalMember
}

func (c *CustomerPortalMembersAddMemberResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalMembersAddMemberResponse) GetCustomerPortalMember() *components.CustomerPortalMember {
	if c == nil {
		return nil
	}
	return c.CustomerPortalMember
}
