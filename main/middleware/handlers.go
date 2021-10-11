package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/koljan1337/second_try/main/models"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
)

type response struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func connect() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := "user=postgres dbname=db123 password=1234 host=localhost sslmode=diable"
	db, err := sql.Open("postgers", os.Getenv(connStr))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
	return db
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var person models.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatalf("Unable to decode request body. %v", err)
	}

	insertID := insertPerson(person)
	res := response{
		ID:      string(rune(insertID)),
		Message: "User created successfully",
	}
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

	res := response{
		ID:      string(rune(int(id))),
		Message: msg,
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
		ID:      string(rune(int(id))),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

//Create person
func insertPerson(person models.Person) int {
	db := connect()

	defer db.Close()

	sqlStatement := `INSERT INTO person (FirstName, LastName, Email, BirthDate, Address, Gender) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int
	err := db.QueryRow(sqlStatement, person.FirstName, person.LastName, person.Email, person.BirthDate, person.Address, person.Gender).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the querry. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

//Get person by ID
func getPerson(id int) (models.Person, error) {
	db := connect()

	defer db.Close()

	var person models.Person

	sqlStatement := `SELECT * FROM person WHERE id=$7`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Email, &person.BirthDate, &person.Address, &person.Gender)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No persons found")
		return person, nil
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
		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Email, &person.BirthDate, &person.Address, &person.Gender)
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

	sqlStatement := `UPDATE person SET FirstName=$1, LastName=$2, Email=$3, BirthDate=$4, Address=$5, Gender=$6 WHERE id=&7`

	res, err := db.Exec(sqlStatement, id, person.FirstName, person.LastName, person.Email, person.BirthDate, person.Address, person.Gender)
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

	sqlStatement := `DELETE FROM person WHERE id=$7`
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
