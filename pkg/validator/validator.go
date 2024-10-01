package validator

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	ValidationErrors map[string]string
	ValidationTag    string
)

const (
	Form ValidationTag = "form"
	JSON ValidationTag = "json"
)

var validate *validator.Validate

// Validate validating request body or other value that must be struct type
func Validate(value interface{}, tag ValidationTag) ValidationErrors {
	if validate == nil {
		validate = validator.New()
	}

	if tag != Form && tag != JSON {
		tag = JSON
	}

	if reflect.ValueOf(value).Kind() != reflect.Struct {
		value = struct{}{}
	}

	// custom validator
	validate.RegisterValidation("dateonly", ValidateDateOnly)

	// register tag to be validated
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get(string(tag)), ",", 2)[0]

		// skip tag if want to be ignored
		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(value)

	if err != nil {
		return formatErrorValidation(err)
	}

	return nil
}

// formatErrorValidation formatting error from go-validator for more readable and understandable error
func formatErrorValidation(err error) ValidationErrors {
	errFields := make(ValidationErrors)

	// make error message for each invalid field
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required", "required_if":
			errFields[err.Field()] = "this field is required"
		case "email":
			errFields[err.Field()] = "invalid email format"
		case "min":
			errFields[err.Field()] = fmt.Sprintf("min length %s characters", err.Param())
		case "max":
			errFields[err.Field()] = fmt.Sprintf("max length %s characters", err.Param())
		case "numeric":
			errFields[err.Field()] = "value must be numeric format"
		case "oneof":
			values := strings.ReplaceAll(err.Param(), " ", ", ")
			errFields[err.Field()] = fmt.Sprintf("must be one of %s", values)
		case "uppercase":
			errFields[err.Field()] = "value must be uppercase"
		case "dateonly":
			errFields[err.Field()] = "value must be date only format"
		default:
			errFields[err.Field()] = err.Error()
		}
	}

	return errFields
}

func ValidateDateOnly(fl validator.FieldLevel) bool {
	val := fl.Field().String()

	_, err := time.Parse(time.DateOnly, val)

	return err == nil
}
