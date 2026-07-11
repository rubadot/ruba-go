package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type BenefitGrantLicenseKeysProperties struct {
	UserProvidedKey *string `json:"user_provided_key,omitempty"`
	LicenseKeyID    *string `json:"license_key_id,omitempty"`
	DisplayKey      *string `json:"display_key,omitempty"`
}

func (b BenefitGrantLicenseKeysProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitGrantLicenseKeysProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (b *BenefitGrantLicenseKeysProperties) GetUserProvidedKey() *string {
	if b == nil {
		return nil
	}
	return b.UserProvidedKey
}

func (b *BenefitGrantLicenseKeysProperties) GetLicenseKeyID() *string {
	if b == nil {
		return nil
	}
	return b.LicenseKeyID
}

func (b *BenefitGrantLicenseKeysProperties) GetDisplayKey() *string {
	if b == nil {
		return nil
	}
	return b.DisplayKey
}
