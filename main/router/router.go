package router

import (
	"github.com/gorilla/mux"
	"github.com/koljan1337/second_try/main/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/persons", middleware.CreatePerson).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/persons", middleware.GetAllPersons).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/persons/{id}", middleware.GetPerson).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/persons/{id}", middleware.UpdatePerson).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/persons/{id}", middleware.DeletePerson).Methods("DELETE", "OPTIONS")

	return router
}
