package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)

func FormatValidationErrors(errs validator.ValidationErrors) map[string]string {
	fmt.Println("formatting")
	formattedErrors := make(map[string]string)
	for _, err := range errs {
		formattedErrors[err.Field()] = fmt.Sprintf("Field is %s", err.Tag())
	}
	fmt.Println(formattedErrors)
	return formattedErrors
}
