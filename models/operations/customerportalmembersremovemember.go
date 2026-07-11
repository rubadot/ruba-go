package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalMembersRemoveMemberRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (c *CustomerPortalMembersRemoveMemberRequest) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

type CustomerPortalMembersRemoveMemberResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (c *CustomerPortalMembersRemoveMemberResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}
