package validator

import (
	"fmt"
	"strings"

	"github.com/sunsetsavorer/grind/internal/exceptions"

	"github.com/go-playground/validator/v10"
)

var validatorMessages = map[string]string{
	"required": "Поле является обязательным",
	"number":   "Значение поля должно быть числом",
	"min":      "Значение поля меньше минимального",
	"max":      "Значение поля превышает допустимое значение",
	"datetime": "Неправильный формат даты",
	"oneof":    "Недопустимое значение",
	"woSpaces": "Значение не должно содержать пробелов",
}

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {

	v := validator.New()

	v.RegisterValidation("woSpaces", woSpaces)

	return &Validator{
		v,
	}
}

func (v Validator) Struct(structInput interface{}) *exceptions.ValidationError {

	err := v.validate.Struct(structInput)
	if err != nil {
		return exceptions.NewValidationError(
			v.getErrors(err.(validator.ValidationErrors)),
		)
	}

	return nil
}

func (v Validator) getErrors(errors validator.ValidationErrors) []exceptions.ValidationField {

	var validationErr []exceptions.ValidationField

	for _, err := range errors {

		field := strings.ToLower(err.Field())
		tag := err.Tag()

		msg, ok := validatorMessages[tag]
		if !ok {
			msg = "Недопустимое значение"
		}

		validationErr = append(validationErr, exceptions.ValidationField{
			Name: field,
			Err:  fmt.Errorf("%s", msg),
		})
	}

	return validationErr
}
