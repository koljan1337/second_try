package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/koljan1337/second_try/main/models"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	_ "os"
	"strconv"
)

type response struct {
	Message string         `json:"message,omitempty"`
	Person  *models.Person `json:"person,omitempty"`
	Msg     string         `json:"msg,omitempty"`
}

func connect() *sql.DB {
	connStr := "user=postgres dbname=database password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	return db
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var person models.Person

	//var loh []validator.FieldError

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatalf("Unable to decode request body. %v", err)
	}

	v := validator.New()
	validationErr := v.Struct(person)

	var res response
	if validationErr != nil {
		res.Message = validationErr.Error()

	}

	if res.Message == "" {
		insertID := insertPerson(person)
		person.ID = insertID

		res = response{
			Person: &person,
		}
	}

	fmt.Println(res)
	json.NewEncoder(w).Encode(res)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	person, err := getPerson(int(id))
	if err != nil {
		log.Fatalf("Unable to get person. %v", err)
	}

	json.NewEncoder(w).Encode(person)
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	persons, err := getAllPersons()
	if err != nil {
		log.Fatalf("Unable to get all persons. %v", err)
	}

	json.NewEncoder(w).Encode(persons)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string to int. %v", err)
	}

	var person models.Person
	err = json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		log.Fatalf("Unable to decode request body. %v", err)
	}

	updatedRows := updatePerson(int(id), person)
	msg := fmt.Sprintf("Person updates successfully. Total affected fields : %v", updatedRows)
	person.ID = id
	res := response{
		Message: "",
		Person:  &person,
		Msg:     msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to covert string into int. %v", err)
	}
	deletedRows := deletePerson(int(id))
	msg := fmt.Sprintf("User deleted successfully. %v", deletedRows)

	res := response{
		Message: "",
		Msg:     msg,
		//Person: ,
	}

	json.NewEncoder(w).Encode(res)
}

//Create person
func insertPerson(person models.Person) int {
	db := connect()

	defer db.Close()

	//if person.Gender != "Male" || person.Gender != "Female" {
	//	log.Fatal("Gender out of range.")
	//	db.Close()
	//}
	var id int

	sqlStatement := `INSERT INTO person (first_name, last_name, email, birth_date, address, gender) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(sqlStatement, person.FirstName, person.LastName, person.Email, person.BirthDate, person.Address, person.Gender).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)

	return id
	//sqlStatement := `INSERT INTO person (first_name, last_name, email, birth_date, address, gender) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	//var id int

	//err := person.bebra
	//if err != nil {
	//	return 0, err
	//}

	//err = db.QueryRow(sqlStatement, person.FirstName, person.LastName, person.Email, person.BirthDate, person.Address, person.Gender).Scan(&id)
}

//Get person by ID
func getPerson(id int) (models.Person, error) {
	db := connect()

	defer db.Close()

	var person models.Person

	sqlStatement := `SELECT * FROM person WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&person.FirstName, &person.LastName, &person.Email, &person.BirthDate, &person.Address, &person.Gender, &person.ID)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No persons found")
		return person, err
	case nil:
		return person, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return person, err
}

//Get all Persons
func getAllPersons() ([]models.Person, error) {
	db := connect()

	defer db.Close()

	var persons []models.Person

	sqlStatement := `SELECT * FROM person`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var person models.Person
		err = rows.Scan(&person.FirstName, &person.LastName, &person.Email, &person.BirthDate, &person.Address, &person.Gender, &person.ID)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		persons = append(persons, person)
	}
	return persons, err
}

//Update Person
func updatePerson(id int, person models.Person) int {
	db := connect()

	defer db.Close()

	//if person.Gender != "Male" || person.Gender != "Female" {
	//	log.Fatal("Gender is out of range.")
	//	db.Close()
	//}

	sqlStatement := `UPDATE person SET first_name=$1, last_name=$2, email=$3, birth_date=$4, address=$5, gender=$6 WHERE id=$7`

	res, err := db.Exec(sqlStatement, person.FirstName, person.LastName, person.Email, person.BirthDate, person.Address, person.Gender, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking affected rows. %v", err)
	}

	fmt.Printf("Total fields affected %v", rowsAffected)

	return int(rowsAffected)
}

//Delete person
func deletePerson(id int) int {
	db := connect()

	defer db.Close()

	sqlStatement := `DELETE FROM person WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the querry. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking affected rows. %v", err)
	}

	fmt.Printf("Total fields affected %v", rowsAffected)

	return int(rowsAffected)
}
