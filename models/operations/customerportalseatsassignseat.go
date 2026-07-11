package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalSeatsAssignSeatSecurity struct {
	CustomerSession *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_customer_session"`
	MemberSession   *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_member_session"`
}

func (c *CustomerPortalSeatsAssignSeatSecurity) GetCustomerSession() *string {
	if c == nil {
		return nil
	}
	return c.CustomerSession
}

func (c *CustomerPortalSeatsAssignSeatSecurity) GetMemberSession() *string {
	if c == nil {
		return nil
	}
	return c.MemberSession
}

type CustomerPortalSeatsAssignSeatResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	CustomerSeat *components.CustomerSeat
}

func (c *CustomerPortalSeatsAssignSeatResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalSeatsAssignSeatResponse) GetCustomerSeat() *components.CustomerSeat {
	if c == nil {
		return nil
	}
	return c.CustomerSeat
}
