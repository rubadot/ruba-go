package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
	"github.com/Rubadot/ruba-go/types"
)

type CountAggregation struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	func_ *string `const:"count" json:"func"`
}

func (c CountAggregation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CountAggregation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *CountAggregation) GetFunc() *string {
	return types.Pointer("count")
}
