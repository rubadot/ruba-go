package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

type CustomFieldCheckboxProperties struct {
	FormLabel       *string `json:"form_label,omitempty"`
	FormHelpText    *string `json:"form_help_text,omitempty"`
	FormPlaceholder *string `json:"form_placeholder,omitempty"`
}

func (c CustomFieldCheckboxProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomFieldCheckboxProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CustomFieldCheckboxProperties) GetFormLabel() *string {
	if c == nil {
		return nil
	}
	return c.FormLabel
}

func (c *CustomFieldCheckboxProperties) GetFormHelpText() *string {
	if c == nil {
		return nil
	}
	return c.FormHelpText
}

func (c *CustomFieldCheckboxProperties) GetFormPlaceholder() *string {
	if c == nil {
		return nil
	}
	return c.FormPlaceholder
}
