package validator

import (
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
)

type ValidationErrorMessageMap map[string]func(validator2.FieldError) string

const (
	requiredMsg  string = "Must be not empty"
	emailMsg            = "Invalid email"
	gteMsg              = "Must be greater than or equal %v"
	eqfieldMsg          = "Does not match to %v"
	nefieldMsg          = "Must be not match to %v"
	minIntMsg           = "Must be greater than or equal %v"
	minStrLenMsg        = "Min length of string is %v"
	minArrLenMsg        = "Min length of array is %v"
	maxIntMsg           = "Must be smaller than or equal %v"
	maxStrLenMsg        = "Max length of string is %v"
	maxArrLenMsg        = "Max length of array is %v"
	mongodbMsg          = "Invalid ObjectId"
	hexcolorMsg         = "Invalid hex color string"
	urlMsg              = "Invalid URL format"
)

var msgMap ValidationErrorMessageMap = map[string]func(validator2.FieldError) string{
	"required": func(err validator2.FieldError) string { return requiredMsg },
	"email":    func(err validator2.FieldError) string { return emailMsg },
	"gte":      func(err validator2.FieldError) string { return fmt.Sprintf(gteMsg, err.Param()) },
	"eqfield":  func(err validator2.FieldError) string { return fmt.Sprintf(eqfieldMsg, err.Param()) },
	"nefield":  func(err validator2.FieldError) string { return fmt.Sprintf(nefieldMsg, err.Param()) },
	"min": func(err validator2.FieldError) string {
		if _, ok := err.Value().(int); ok {
			return fmt.Sprintf(minIntMsg, err.Param())
		} else if _, ok := err.Value().(string); ok {
			return fmt.Sprintf(minStrLenMsg, err.Param())
		}
		return fmt.Sprintf(minArrLenMsg, err.Param())
	},
	"max": func(err validator2.FieldError) string {
		if _, ok := err.Value().(int); ok {
			return fmt.Sprintf(maxIntMsg, err.Param())
		} else if _, ok := err.Value().(string); ok {
			return fmt.Sprintf(maxStrLenMsg, err.Param())
		}
		return fmt.Sprintf(maxArrLenMsg, err.Param())
	},
	"mongodb":  func(err validator2.FieldError) string { return mongodbMsg },
	"hexcolor": func(err validator2.FieldError) string { return hexcolorMsg },
	"url":      func(err validator2.FieldError) string { return urlMsg },
}

func convert(err validator2.FieldError) *ValidationError {
	res := &ValidationError{Field: err.Field(), Value: err.Value()}
	res.Message = msgMap[err.ActualTag()](err)

	return res
}
