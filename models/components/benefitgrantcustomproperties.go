package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type BenefitGrantCustomProperties struct {
}

func (b BenefitGrantCustomProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitGrantCustomProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, nil); err != nil {
		return err
	}
	return nil
}
