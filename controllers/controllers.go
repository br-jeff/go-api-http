package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/br-jeff/go-api-http/models"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Route")
}

func ListAllPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.People)
}
