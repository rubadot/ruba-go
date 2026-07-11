package rubago

import (
	"github.com/Rubadot/ruba-go/internal/config"
	"github.com/Rubadot/ruba-go/internal/hooks"
)

type CustomerPortal struct {
	BenefitGrants   *RubaBenefitGrants
	Customers       *RubaCustomers
	CustomerMeters  *RubaCustomerMeters
	Seats           *Seats
	CustomerSession *CustomerSession
	Downloadables   *Downloadables
	LicenseKeys     *RubaLicenseKeys
	Members         *RubaMembers
	Orders          *RubaOrders
	Organizations   *RubaOrganizations
	Subscriptions   *RubaSubscriptions
	Wallets         *Wallets

	rootSDK          *Ruba
	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

func newCustomerPortal(rootSDK *Ruba, sdkConfig config.SDKConfiguration, hooks *hooks.Hooks) *CustomerPortal {
	return &CustomerPortal{
		rootSDK:          rootSDK,
		sdkConfiguration: sdkConfig,
		hooks:            hooks,
		BenefitGrants:    newRubaBenefitGrants(rootSDK, sdkConfig, hooks),
		Customers:        newRubaCustomers(rootSDK, sdkConfig, hooks),
		CustomerMeters:   newRubaCustomerMeters(rootSDK, sdkConfig, hooks),
		Seats:            newSeats(rootSDK, sdkConfig, hooks),
		CustomerSession:  newCustomerSession(rootSDK, sdkConfig, hooks),
		Downloadables:    newDownloadables(rootSDK, sdkConfig, hooks),
		LicenseKeys:      newRubaLicenseKeys(rootSDK, sdkConfig, hooks),
		Members:          newRubaMembers(rootSDK, sdkConfig, hooks),
		Orders:           newRubaOrders(rootSDK, sdkConfig, hooks),
		Organizations:    newRubaOrganizations(rootSDK, sdkConfig, hooks),
		Subscriptions:    newRubaSubscriptions(rootSDK, sdkConfig, hooks),
		Wallets:          newWallets(rootSDK, sdkConfig, hooks),
	}
}
