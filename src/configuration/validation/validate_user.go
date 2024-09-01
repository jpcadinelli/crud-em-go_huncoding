package validation

import (
	"crud-em-go_huncoding/src/configuration/rest_err"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *rest_err {
	var (
		jsonErr             *json.UnmarshalTypeError
		jsonValidationError validator.ValidationErrors
	)

	if errors.As(validationErr, jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, jsonValidationError) {
		errosCauses := []rest_err.Causes{}

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}
			errosCauses = append(errosCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errosCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
