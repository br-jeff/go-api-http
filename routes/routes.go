package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/br-jeff/go-api-http/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.IndexController)
	r.HandleFunc("/api/person", controllers.ListAllPeople).Methods("Get")
	r.HandleFunc("/api/person/{id}", controllers.FindPersonById).Methods("Get")

	log.Fatal(http.ListenAndServe(":8080", r))
}
