package errors

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")

	ErrNotFound = errors.New("your requested Item is not found")

	ErrConflict = errors.New("your Item already exist")

	ErrBadParamInput = errors.New("given Param is not valid")
)

type FieldError struct {
	Msg string `json:"msg"`
	Tag string `json:"tag"`
}

type ValidationError struct {
	Errors map[string]FieldError `json:"errors"`
}
