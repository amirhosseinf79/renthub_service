package pkg

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func ValidateRequestBody[T any](request *T, c fiber.Ctx) (map[string]string, error) {
	if err := c.Bind().Body(request); err != nil {
		return map[string]string{
			"error": "Invalid request data",
		}, errors.New("invalid request data")
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Field '%s' failed validation with tag '%s'", err.Field(), err.Tag())
		}
		return errors, err
	}
	return nil, nil
}

func ValidateQueryParams[T any](request *T, c fiber.Ctx) (map[string]string, error) {
	if err := c.Bind().Query(request); err != nil {
		return map[string]string{
			"error": "Invalid request data",
		}, errors.New("invalid request data")
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("Field '%s' failed validation with tag '%s'", err.Field(), err.Tag())
		}
		return errors, err
	}
	return nil, nil
}
