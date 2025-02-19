package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/wrferreira1003/concorrencia-go-leilao/config/rest_err.go"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTrans := ut.New(en, en)
		transl, _ = enTrans.GetTranslator("en")
		enTranslations.RegisterDefaultTranslations(v, transl)
	}
}

// ValidateErrors validates the validation errors
func ValidateErrors(validationErrors error) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validationErrors, &jsonErr) {
		return rest_err.NewBadRequestError("invalid field type")
	} else if errors.As(validationErrors, &jsonValidation) {
		errorCauses := []rest_err.Causes{}
		for _, e := range validationErrors.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			})
		}

		return rest_err.NewBadRequestError("invalid request", errorCauses...)
	} else {
		return rest_err.NewBadRequestError("invalid request")
	}

}
