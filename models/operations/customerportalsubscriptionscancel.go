package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalSubscriptionsCancelSecurity struct {
	CustomerSession *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_customer_session"`
	MemberSession   *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_member_session"`
}

func (c *CustomerPortalSubscriptionsCancelSecurity) GetCustomerSession() *string {
	if c == nil {
		return nil
	}
	return c.CustomerSession
}

func (c *CustomerPortalSubscriptionsCancelSecurity) GetMemberSession() *string {
	if c == nil {
		return nil
	}
	return c.MemberSession
}

type CustomerPortalSubscriptionsCancelRequest struct {
	// The subscription ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (c *CustomerPortalSubscriptionsCancelRequest) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

type CustomerPortalSubscriptionsCancelResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Customer subscription is canceled.
	CustomerSubscription *components.CustomerSubscription
}

func (c *CustomerPortalSubscriptionsCancelResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalSubscriptionsCancelResponse) GetCustomerSubscription() *components.CustomerSubscription {
	if c == nil {
		return nil
	}
	return c.CustomerSubscription
}
