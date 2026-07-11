package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalCustomersUpdateSecurity struct {
	CustomerSession *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_customer_session"`
	MemberSession   *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_member_session"`
}

func (c *CustomerPortalCustomersUpdateSecurity) GetCustomerSession() *string {
	if c == nil {
		return nil
	}
	return c.CustomerSession
}

func (c *CustomerPortalCustomersUpdateSecurity) GetMemberSession() *string {
	if c == nil {
		return nil
	}
	return c.MemberSession
}

type CustomerPortalCustomersUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Customer updated.
	CustomerPortalCustomer *components.CustomerPortalCustomer
}

func (c *CustomerPortalCustomersUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalCustomersUpdateResponse) GetCustomerPortalCustomer() *components.CustomerPortalCustomer {
	if c == nil {
		return nil
	}
	return c.CustomerPortalCustomer
}
