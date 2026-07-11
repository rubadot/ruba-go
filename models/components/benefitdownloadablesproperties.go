package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type BenefitDownloadablesProperties struct {
	Archived map[string]bool `json:"archived"`
	Files    []string        `json:"files"`
}

func (b BenefitDownloadablesProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitDownloadablesProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, []string{"archived", "files"}); err != nil {
		return err
	}
	return nil
}

func (b *BenefitDownloadablesProperties) GetArchived() map[string]bool {
	if b == nil {
		return map[string]bool{}
	}
	return b.Archived
}

func (b *BenefitDownloadablesProperties) GetFiles() []string {
	if b == nil {
		return []string{}
	}
	return b.Files
}
