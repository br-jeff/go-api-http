package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/br-jeff/go-api-http/database"
	"github.com/br-jeff/go-api-http/models"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Route")
}

func ListAllPeople(w http.ResponseWriter, r *http.Request) {
	var p []models.Person

	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func FindPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person

	database.DB.First(&person, id)
	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Create(&person)
	json.NewEncoder(w).Encode(person)
}
