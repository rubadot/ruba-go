package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type BenefitGrantDownloadablesProperties struct {
	Files []string `json:"files,omitempty"`
}

func (b BenefitGrantDownloadablesProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitGrantDownloadablesProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (b *BenefitGrantDownloadablesProperties) GetFiles() []string {
	if b == nil {
		return nil
	}
	return b.Files
}
