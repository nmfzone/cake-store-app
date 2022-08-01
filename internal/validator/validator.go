package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/nmfzone/privy-cake-store/internal/errors"
	"github.com/nmfzone/privy-cake-store/internal/utils"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strings"
)

type RequestValidator struct {
	validator *validator.Validate
}

var Engine *validator.Validate
var Translator ut.Translator

func init() {
	Engine = binding.Validator.Engine().(*validator.Validate)

	RegisterNullableValidator(Engine)
	RegisterCustomValidator(Engine)

	// re-format StructFields name, so it can be customized later easily
	Engine.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return ":" + fld.Name + ":"
	})

	english := en.New()
	uni := ut.New(english, english)
	Translator, _ = uni.GetTranslator("en")
	transErr := entranslations.RegisterDefaultTranslations(Engine, Translator)

	if transErr != nil {
		errMsg := fmt.Sprintf("RegisterDefaultTranslations Error: %v", transErr)
		defer log.Fatal().Msg(errMsg)
		panic(errMsg)
	}
}

func NewRequestValidator() RequestValidator {
	return RequestValidator{}
}

func (r *RequestValidator) Validate(s interface{}) (bool, errors.ValidationError) {
	if reflect.ValueOf(s).Kind() != reflect.Struct {
		errMsg := fmt.Sprintf(
			"Validator value not supported, because %v is not struct type",
			reflect.TypeOf(s))
		defer log.Fatal().Msg(errMsg)
		panic(errMsg)
	}

	err := Engine.Struct(s)

	return NewValidationError(s, err)
}

func ValidateSingleField(value string, tag string, field string) errors.ValidationError {
	vErr := errors.ValidationError{}

	rules := map[string]string{
		"GenericField": tag,
	}

	data := genericValidatorData{
		GenericField: value,
	}

	Engine.RegisterStructValidationMapRules(rules, genericValidatorData{})

	err := Engine.Struct(data)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		_, vErr = NewValidationError(data, errs)

		fErr := vErr.Errors["genericfield"]
		delete(vErr.Errors, "genericfield")

		title := cases.Title(language.English)
		fErr.Msg = strings.Replace(fErr.Msg, "GenericField", title.String(field), 1)
		vErr.Errors[field] = fErr
	}

	return vErr
}

func NewValidationError(s interface{}, err error) (bool, errors.ValidationError) {
	res := errors.ValidationError{}

	if err == nil {
		return true, res
	}

	res.Errors = make(map[string]errors.FieldError)
	errs := err.(validator.ValidationErrors)

	for _, v := range errs {
		meta := utils.GetStructFieldMetadata(s, v.StructField())

		errResult := errors.FieldError{
			Msg: v.Translate(Translator),
			Tag: v.Tag(),
		}

		key := v.StructField()
		mKey := key

		if meta != nil {
			fv := meta.Get("form")

			if fv != "" {
				key = fv

				title := cases.Title(language.English)
				mKey = title.String(strings.ReplaceAll(fv, "_", " "))
			}
		}

		errResult.Msg = strings.Replace(errResult.Msg, v.Field(), mKey, 1)

		res.Errors[strings.ToLower(key)] = errResult
	}

	return res.Errors == nil, res
}
