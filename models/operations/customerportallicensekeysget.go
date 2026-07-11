package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomerPortalLicenseKeysGetSecurity struct {
	CustomerSession *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_customer_session"`
	MemberSession   *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_member_session"`
}

func (c *CustomerPortalLicenseKeysGetSecurity) GetCustomerSession() *string {
	if c == nil {
		return nil
	}
	return c.CustomerSession
}

func (c *CustomerPortalLicenseKeysGetSecurity) GetMemberSession() *string {
	if c == nil {
		return nil
	}
	return c.MemberSession
}

type CustomerPortalLicenseKeysGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (c *CustomerPortalLicenseKeysGetRequest) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

type CustomerPortalLicenseKeysGetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	LicenseKeyWithActivations *components.LicenseKeyWithActivations
}

func (c *CustomerPortalLicenseKeysGetResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}

func (c *CustomerPortalLicenseKeysGetResponse) GetLicenseKeyWithActivations() *components.LicenseKeyWithActivations {
	if c == nil {
		return nil
	}
	return c.LicenseKeyWithActivations
}
