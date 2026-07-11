package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

// BenefitCustomProperties - Properties for a benefit of type `custom`.
type BenefitCustomProperties struct {
	Note *string `json:"note"`
}

func (b BenefitCustomProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitCustomProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (b *BenefitCustomProperties) GetNote() *string {
	if b == nil {
		return nil
	}
	return b.Note
}
