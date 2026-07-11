package apierrors

import (
	"encoding/json"
)

type NotPermitted struct {
	Error_ string `const:"NotPermitted" json:"error"`
	Detail string `json:"detail"`
}

var _ error = &NotPermitted{}

func (e *NotPermitted) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
