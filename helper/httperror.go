package helper

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type apiError struct {
	Param   string
	Message string
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	default:
		return fe.Error()
	}
}

func HandleValidatorErr(err error) []apiError {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]apiError, len(ve))
		for i, fe := range ve {
			out[i] = apiError{fe.Field(), msgForTag(fe)}
		}
		return out
	}

	return nil
}
