package errors

import (
	"errors"
	"strings"

	"github.com/codercollo/blog/internal/providers/validation"
	"github.com/go-playground/validator/v10"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFromErrors(err error) {
	var validatonErrors validator.ValidationErrors

	if errors.As(err, &validatonErrors) {
		for _, fieldError := range validatonErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))

		}
	}
}

func Add(key string, value string) {
	errorsList[strings.ToLower(key)] = value
}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorsList
}
