package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type BenefitGrantFeatureFlagProperties struct {
}

func (b BenefitGrantFeatureFlagProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitGrantFeatureFlagProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, nil); err != nil {
		return err
	}
	return nil
}
