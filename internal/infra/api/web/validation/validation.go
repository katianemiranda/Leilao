package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	rest_error "github.com/katianemiranda/leilao/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTransl := ut.New(en, en)
		transl, _ = enTransl.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(value, transl)

	}
}

func ValidateErr(validation_err error) *rest_error.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors
	if errors.As(validation_err, &jsonErr) {
		return rest_error.NewNotFoundError("Invalid JSON")
	} else if errors.As(validation_err, &jsonValidation) {
		errorCauses := []rest_error.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_error.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			})
		}
		return rest_error.NewBadRequestError("Unknown error", errorCauses...)
	} else {
		return rest_error.NewBadRequestError("Unknown error")
	}

}
