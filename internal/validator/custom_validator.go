package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type genericValidatorData struct {
	GenericField string
}

func RegisterCustomValidator(engine *validator.Validate) {
	var err error = nil

	err = engine.RegisterValidation("maxx", ValidateMaxx)

	if err != nil {
		panic(err)
	}
}

func ValidateMaxx(fl validator.FieldLevel) bool {
	log.Info().Msg(fl.FieldName())
	log.Info().Msg(fl.Field().String())
	return true
}
