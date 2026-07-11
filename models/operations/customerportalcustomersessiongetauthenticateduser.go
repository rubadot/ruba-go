package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalCustomerSessionGetAuthenticatedUserSecurity struct {
	CustomerSession *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_customer_session"`
	MemberSession   *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_member_session"`
}

func (c *CustomerPortalCustomerSessionGetAuthenticatedUserSecurity) GetCustomerSession() *string {
	if c == nil {
		return nil
	}
	return c.CustomerSession
}

func (c *CustomerPortalCustomerSessionGetAuthenticatedUserSecurity) GetMemberSession() *string {
	if c == nil {
		return nil
	}
	return c.MemberSession
}

type CustomerPortalCustomerSessionGetAuthenticatedUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	PortalAuthenticatedUser *components.PortalAuthenticatedUser
}

func (c *CustomerPortalCustomerSessionGetAuthenticatedUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalCustomerSessionGetAuthenticatedUserResponse) GetPortalAuthenticatedUser() *components.PortalAuthenticatedUser {
	if c == nil {
		return nil
	}
	return c.PortalAuthenticatedUser
}
