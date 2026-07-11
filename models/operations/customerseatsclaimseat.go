package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerSeatsClaimSeatResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	CustomerSeatClaimResponse *components.CustomerSeatClaimResponse
}

func (c *CustomerSeatsClaimSeatResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerSeatsClaimSeatResponse) GetCustomerSeatClaimResponse() *components.CustomerSeatClaimResponse {
	if c == nil {
		return nil
	}
	return c.CustomerSeatClaimResponse
}
