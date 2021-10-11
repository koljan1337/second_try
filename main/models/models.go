package models

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	ID        int    `json:"id"`
}
