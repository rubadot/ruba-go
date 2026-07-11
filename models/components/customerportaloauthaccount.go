package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type CustomerPortalOAuthAccount struct {
	AccountID       string  `json:"account_id"`
	AccountUsername *string `json:"account_username"`
}

func (c CustomerPortalOAuthAccount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomerPortalOAuthAccount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, []string{"account_id"}); err != nil {
		return err
	}
	return nil
}

func (c *CustomerPortalOAuthAccount) GetAccountID() string {
	if c == nil {
		return ""
	}
	return c.AccountID
}

func (c *CustomerPortalOAuthAccount) GetAccountUsername() *string {
	if c == nil {
		return nil
	}
	return c.AccountUsername
}
