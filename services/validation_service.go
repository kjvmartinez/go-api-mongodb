package services

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate struct
func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}