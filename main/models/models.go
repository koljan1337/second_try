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
	Gender    string `json:"gender"`
	ID        int    `json:"id"`
}

//func main() {
//	v := validator.New()
//	person := Person{Gender: "person"}
//	err := v.Struct(person)
//	for _, e := range err.(validator.ValidationErrors) {
//		fmt.Println(e)
//	}
//}

//var validate *validator.Validate
//
//func init() {
//	validate = validator.New()
//}
//
//func (person *Person) Validate() error {
//	if err := validate.Struct(person); err != nil {
//		if _, ok := err.(*validator.InvalidValidationError); ok {
//			return err
//		}
//		return err
//	}
//	return nil
//}

//func (person Person) IsValid() (errs url.Values) {
//
//	if person.FirstName == "" {
//		errs.Add("name", "Name is required filed")
//	}
//
//	if len(person.FirstName) > 100 {
//		errs.Add("name", "Name must be a maximum of 100 characters in length")
//	}
//
//	if person.LastName == "" {
//		errs.Add("last_name", "Last name is required filed")
//	}
//
//	if len(person.LastName) > 100 {
//		errs.Add("last_name", "Last name must be a maximum of 100 characters in length")
//	}
//
//	if person.Email == "" {
//		errs.Add("email", "Email is required filed")
//	}
//
//	if len(person.Email) > 200 {
//		errs.Add("email", "email must be a maximum of 100 characters in length")
//	}
//
//
//	return errs
//}

//func (person Person) IsValidSurname() (errs url.Values) {
//
//
//	return errs
//}
