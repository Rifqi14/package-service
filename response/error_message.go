package response

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gitlab.com/shoesmart2.1/backend/packages/str"
	"reflect"
	"strings"
)

type IErrorMessage interface {
	CreateError() ErrorResponseVm
}

type ErrorFromValidator struct {
	ErrorMessage validator.ValidationErrors
	Translator   ut.Translator
}

func newErrorFromValidator(errMessage validator.ValidationErrors, translator ut.Translator) IErrorMessage {
	return &ErrorFromValidator{ErrorMessage: errMessage, Translator: translator}
}

func (errorResponse ErrorFromValidator) CreateError() ErrorResponseVm {

	errorMessage := map[string][]string{}
	errorTranslation := errorResponse.ErrorMessage.Translate(errorResponse.Translator)

	reflect.TypeOf(errorMessage).String()

	for _, err := range errorResponse.ErrorMessage {
		errKey := str.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
		)
	}

	return newErrorResponse(errorMessage)
}

type BasicError struct {
	ErrorMessage error
}

func newBasicError(errMessage error) IErrorMessage {
	return &BasicError{ErrorMessage: errMessage}
}

func (errorResponse BasicError) CreateError() ErrorResponseVm {
	return newErrorResponse(errorResponse.ErrorMessage.Error())
}
