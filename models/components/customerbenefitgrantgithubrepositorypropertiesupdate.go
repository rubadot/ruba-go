package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type CustomerBenefitGrantGitHubRepositoryPropertiesUpdate struct {
	AccountID *string `json:"account_id"`
}

func (c CustomerBenefitGrantGitHubRepositoryPropertiesUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomerBenefitGrantGitHubRepositoryPropertiesUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CustomerBenefitGrantGitHubRepositoryPropertiesUpdate) GetAccountID() *string {
	if c == nil {
		return nil
	}
	return c.AccountID
}
