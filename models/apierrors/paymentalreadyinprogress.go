package apierrors

import (
	"encoding/json"
)

type PaymentAlreadyInProgress struct {
	Error_ string `const:"PaymentAlreadyInProgress" json:"error"`
	Detail string `json:"detail"`
}

var _ error = &PaymentAlreadyInProgress{}

func (e *PaymentAlreadyInProgress) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
