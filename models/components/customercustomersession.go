package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
	"time"
)

type CustomerCustomerSession struct {
	ExpiresAt time.Time `json:"expires_at"`
	ReturnURL *string   `json:"return_url"`
}

func (c CustomerCustomerSession) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomerCustomerSession) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CustomerCustomerSession) GetExpiresAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.ExpiresAt
}

func (c *CustomerCustomerSession) GetReturnURL() *string {
	if c == nil {
		return nil
	}
	return c.ReturnURL
}
