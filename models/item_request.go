package models

import "github.com/go-playground/validator/v10"

type ItemRequest struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func (itemInput ItemRequest) ValidateStruct() []*ErrorResponse {
	var errors []*ErrorResponse

	validate := validator.New()
	validationError := validate.Struct(itemInput)

	if validationError != nil {
		for _, err := range validationError.(validator.ValidationErrors) {
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Field = err.Field()
			errors = append(errors, &element)
		}
	}

	return errors
}
