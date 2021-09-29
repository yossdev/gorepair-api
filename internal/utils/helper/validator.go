package helper

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateInputs(dataSet interface{}) (bool, map[string]string) {

	validate := validator.New()

	err := validate.Struct(dataSet)

	if err != nil {
		errors := make(map[string]string)

		reflected := reflect.ValueOf(dataSet)
		
		for _, err := range err.(validator.ValidationErrors) {

			field, _ := reflected.Type().FieldByName(err.StructField())

			var name string

			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = "The " + name + " is required"
			default:
				errors[name] = "The " + name + " is invalid"
			}
		}
		return false, errors
	}
	return true, nil
}
