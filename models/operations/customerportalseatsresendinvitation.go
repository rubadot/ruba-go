package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalSeatsResendInvitationSecurity struct {
	CustomerSession *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_customer_session"`
	MemberSession   *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_member_session"`
}

func (c *CustomerPortalSeatsResendInvitationSecurity) GetCustomerSession() *string {
	if c == nil {
		return nil
	}
	return c.CustomerSession
}

func (c *CustomerPortalSeatsResendInvitationSecurity) GetMemberSession() *string {
	if c == nil {
		return nil
	}
	return c.MemberSession
}

type CustomerPortalSeatsResendInvitationRequest struct {
	SeatID string `pathParam:"style=simple,explode=false,name=seat_id"`
}

func (c *CustomerPortalSeatsResendInvitationRequest) GetSeatID() string {
	if c == nil {
		return ""
	}
	return c.SeatID
}

type CustomerPortalSeatsResendInvitationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	CustomerSeat *components.CustomerSeat
}

func (c *CustomerPortalSeatsResendInvitationResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalSeatsResendInvitationResponse) GetCustomerSeat() *components.CustomerSeat {
	if c == nil {
		return nil
	}
	return c.CustomerSeat
}
