package validator

import (
	"fmt"
)

type ValidationError struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: field %q: %v", e.Field, e.Message)
}
