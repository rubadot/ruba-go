package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerSeatsGetClaimInfoRequest struct {
	InvitationToken string `pathParam:"style=simple,explode=false,name=invitation_token"`
}

func (c *CustomerSeatsGetClaimInfoRequest) GetInvitationToken() string {
	if c == nil {
		return ""
	}
	return c.InvitationToken
}

type CustomerSeatsGetClaimInfoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	SeatClaimInfo *components.SeatClaimInfo
}

func (c *CustomerSeatsGetClaimInfoResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerSeatsGetClaimInfoResponse) GetSeatClaimInfo() *components.SeatClaimInfo {
	if c == nil {
		return nil
	}
	return c.SeatClaimInfo
}
