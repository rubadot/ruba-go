package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type ProductsCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Product created.
	Product *components.Product
}

func (p *ProductsCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if p == nil {
		return components.HTTPMetadata{}
	}
	return p.HTTPMeta
}

func (p *ProductsCreateResponse) GetProduct() *components.Product {
	if p == nil {
		return nil
	}
	return p.Product
}
