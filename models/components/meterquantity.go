package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
	"time"
)

type MeterQuantity struct {
	// The timestamp for the current period.
	Timestamp time.Time `json:"timestamp"`
	// The quantity for the current period.
	Quantity float64 `json:"quantity"`
}

func (m MeterQuantity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeterQuantity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (m *MeterQuantity) GetTimestamp() time.Time {
	if m == nil {
		return time.Time{}
	}
	return m.Timestamp
}

func (m *MeterQuantity) GetQuantity() float64 {
	if m == nil {
		return 0.0
	}
	return m.Quantity
}
