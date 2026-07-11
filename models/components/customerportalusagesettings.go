package components

type CustomerPortalUsageSettings struct {
	Show bool `json:"show"`
}

func (c *CustomerPortalUsageSettings) GetShow() bool {
	if c == nil {
		return false
	}
	return c.Show
}
