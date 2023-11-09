package validator

import (
	validator2 "github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Validator interface {
	Validate(in interface{}) []ValidationError
}

type validator struct {
	validator *validator2.Validate
}

func NewValidator() Validator {
	var v validator
	v.validator = validator2.New()

	// Get json tag value
	v.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return &v
}

func (v *validator) Validate(in interface{}) []ValidationError {
	if errs := v.validator.Struct(in); errs != nil {
		var res []ValidationError
		for _, err := range errs.(validator2.ValidationErrors) {
			res = append(res, *convert(err))
		}

		return res
	}

	return nil
}
