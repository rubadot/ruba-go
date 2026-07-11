package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type RefundsCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Refund created.
	Refund *components.Refund
}

func (r *RefundsCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if r == nil {
		return components.HTTPMetadata{}
	}
	return r.HTTPMeta
}

func (r *RefundsCreateResponse) GetRefund() *components.Refund {
	if r == nil {
		return nil
	}
	return r.Refund
}
