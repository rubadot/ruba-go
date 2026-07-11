package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type CustomFieldsDeleteRequest struct {
	// The custom field ID.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (c *CustomFieldsDeleteRequest) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

type CustomFieldsDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (c *CustomFieldsDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if c == nil {
		return components.HTTPMetadata{}
	}
	return c.HTTPMeta
}
