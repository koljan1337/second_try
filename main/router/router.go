package router

import (
	"github.com/gorilla/mux"
	"github.com/koljan1337/second_try/main/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/persons", middleware.CreatePerson).Methods("POST", "OPTIONS")
	router.HandleFunc("/persons", middleware.GetAllPersons).Methods("GET", "OPTIONS")
	router.HandleFunc("/persons/{id}", middleware.GetPerson).Methods("GET", "OPTIONS")
	router.HandleFunc("/persons/{id}", middleware.UpdatePerson).Methods("PUT", "OPTIONS")
	router.HandleFunc("/persons/{id}", middleware.DeletePerson).Methods("DELETE", "OPTIONS")

	return router
}
