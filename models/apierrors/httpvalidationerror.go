package apierrors

import (
	"encoding/json"
	"github.com/Rubadot/ruba-go/models/components"
)

type HTTPValidationError struct {
	Detail []components.ValidationError `json:"detail,omitempty"`
}

var _ error = &HTTPValidationError{}

func (e *HTTPValidationError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
