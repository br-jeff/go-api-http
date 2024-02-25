package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/br-jeff/go-api-http/models"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Route")
}

func ListAllPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.People)
}

func FindPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, person := range models.People {
		if strconv.Itoa(person.Id) == id {
			json.NewEncoder(w).Encode(person)
		}
	}
}
