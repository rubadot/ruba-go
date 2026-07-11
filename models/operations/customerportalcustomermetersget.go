package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalCustomerMetersGetSecurity struct {
	CustomerSession *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_customer_session"`
	MemberSession   *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_member_session"`
}

func (c *CustomerPortalCustomerMetersGetSecurity) GetCustomerSession() *string {
	if c == nil {
		return nil
	}
	return c.CustomerSession
}

func (c *CustomerPortalCustomerMetersGetSecurity) GetMemberSession() *string {
	if c == nil {
		return nil
	}
	return c.MemberSession
}

type CustomerPortalCustomerMetersGetRequest struct {
	// The customer meter ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (c *CustomerPortalCustomerMetersGetRequest) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

type CustomerPortalCustomerMetersGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	CustomerCustomerMeter *components.CustomerCustomerMeter
}

func (c *CustomerPortalCustomerMetersGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalCustomerMetersGetResponse) GetCustomerCustomerMeter() *components.CustomerCustomerMeter {
	if c == nil {
		return nil
	}
	return c.CustomerCustomerMeter
}
