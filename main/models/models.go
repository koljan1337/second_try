package models

import (
	_ "github.com/go-playground/validator"
)

type Person struct {
	FirstName string `json:"first_name" validate:"required,max=100"`
	LastName  string `json:"last_name" validate:"required,max=100"`
	Email     string `json:"email" validate:"required,email,max=200"`
	BirthDate string `json:"birth_date" validate:"required"`
	Address   string `json:"address" validate:"max=200"`
	Gender    string `json:"gender" validate:"oneof=Male Female"`
	ID        int    `json:"id"`
}
