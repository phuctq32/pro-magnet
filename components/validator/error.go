package validator

import (
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	"strings"
)

type ValidationError struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: field %q: %v", e.Field, e.Message)
}

func convert(err validator2.FieldError) *ValidationError {
	res := &ValidationError{Field: err.Field(), Value: err.Value()}
	switch err.ActualTag() {
	case "required":
		res.Message = "Must be not empty"
	case "email":
		res.Message = "Invalid email"
	case "gte":
		res.Message = fmt.Sprintf("Must be greater than or equal %v", err.Param())
	case "eqfield":
		res.Message = fmt.Sprintf("Does not match to %v", strings.ToLower(err.Param()))
	case "nefield":
		res.Message = fmt.Sprintf("Must be not match to %v", strings.ToLower(err.Param()))
	case "min":
		if _, ok := res.Value.(int); ok {
			res.Message = fmt.Sprintf("Must be greater than or equal %v", err.Param())
		} else if _, ok := res.Value.(string); ok {
			res.Message = fmt.Sprintf("Min length of string is %v", err.Param())
		} else {
			res.Message = fmt.Sprintf("Min length of array is %v", err.Param())
		}
	case "max":
		if _, ok := res.Value.(int); ok {
			res.Message = fmt.Sprintf("Must be smaller than or equal %v", err.Param())
		} else if _, ok := res.Value.(string); ok {
			res.Message = fmt.Sprintf("Max length of string is %v", err.Param())
		} else {
			res.Message = fmt.Sprintf("Max length of array is %v", err.Param())
		}
	case "mongodb":
		res.Message = "Invalid ObjectId"
	case "hexcolor":
		res.Message = "Invalid hex color string"
	}

	return res
}
