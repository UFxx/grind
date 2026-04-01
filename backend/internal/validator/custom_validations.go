package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func woSpaces(fl validator.FieldLevel) bool {

	return !strings.Contains(fl.Field().String(), " ")
}
