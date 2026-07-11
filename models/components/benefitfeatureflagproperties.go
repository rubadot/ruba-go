package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

// BenefitFeatureFlagProperties - Properties for a benefit of type `feature_flag`.
type BenefitFeatureFlagProperties struct {
}

func (b BenefitFeatureFlagProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitFeatureFlagProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, nil); err != nil {
		return err
	}
	return nil
}
