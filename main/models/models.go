package models

import "time"

type Person struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
	Address   string    `json:"address"`
	Gender    string    `json:"gender"`
	ID        int       `json:"id"`
}
