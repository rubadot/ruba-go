package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type BenefitsDeleteRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (b *BenefitsDeleteRequest) GetID() string {
	if b == nil {
		return ""
	}
	return b.ID
}

type BenefitsDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (b *BenefitsDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if b == nil {
		return components.HTTPMetadata{}
	}
	return b.HTTPMeta
}
