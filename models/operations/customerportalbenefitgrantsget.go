package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalBenefitGrantsGetSecurity struct {
	CustomerSession *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_customer_session"`
	MemberSession   *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_member_session"`
}

func (c *CustomerPortalBenefitGrantsGetSecurity) GetCustomerSession() *string {
	if c == nil {
		return nil
	}
	return c.CustomerSession
}

func (c *CustomerPortalBenefitGrantsGetSecurity) GetMemberSession() *string {
	if c == nil {
		return nil
	}
	return c.MemberSession
}

type CustomerPortalBenefitGrantsGetRequest struct {
	// The benefit grant ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (c *CustomerPortalBenefitGrantsGetRequest) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

type CustomerPortalBenefitGrantsGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	CustomerBenefitGrant *components.CustomerBenefitGrant
}

func (c *CustomerPortalBenefitGrantsGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalBenefitGrantsGetResponse) GetCustomerBenefitGrant() *components.CustomerBenefitGrant {
	if c == nil {
		return nil
	}
	return c.CustomerBenefitGrant
}
