package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type BenefitDownloadablesSubscriberProperties struct {
	ActiveFiles []string `json:"active_files"`
}

func (b BenefitDownloadablesSubscriberProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitDownloadablesSubscriberProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, []string{"active_files"}); err != nil {
		return err
	}
	return nil
}

func (b *BenefitDownloadablesSubscriberProperties) GetActiveFiles() []string {
	if b == nil {
		return []string{}
	}
	return b.ActiveFiles
}
