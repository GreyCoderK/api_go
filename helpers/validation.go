package helpers

import (
	"../dtos"
	"../langs"
	"gopkg.in/go-playground/validator.v8"
)

func GenerateValidationResponse(err error) (res dtos.ValidationResponse) {
	res.Success = false
	var validations []dtos.Validation

	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		field, rule := value.Field, value.Tag
		validation := dtos.Validation{
			Field:   field,
			Message: langs.GenerateValidationMessage(field, rule),
		}

		validations = append(validations, validation)
	}

	res.Validations = validations

	return res
}
