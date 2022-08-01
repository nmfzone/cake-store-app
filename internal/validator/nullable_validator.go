package validator

import (
	"github.com/Thor-x86/nullable"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// @TODO: Add cutom type func for another nullable type

func RegisterNullableValidator(engine *validator.Validate) {
	engine.RegisterCustomTypeFunc(
		validateNullableField,
		nullable.Int{},
		nullable.Int64{},
		nullable.String{},
	)
}

func validateNullableField(field reflect.Value) interface{} {
	if val, ok := field.Interface().(nullable.String); ok {
		res := val.Get()
		if res != nil {
			return *res
		}
	} else if val, ok := field.Interface().(nullable.Int); ok {
		res := val.Get()
		if res != nil {
			return *res
		}
	} else if val, ok := field.Interface().(nullable.Int64); ok {
		res := val.Get()
		if res != nil {
			return *res
		}
	}

	return nil
}
