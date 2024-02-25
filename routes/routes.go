package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/br-jeff/go-api-http/controllers"
	"github.com/br-jeff/go-api-http/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.IndexController)
	r.HandleFunc("/api/person", controllers.ListAllPeople).Methods("Get")
	r.HandleFunc("/api/person/{id}", controllers.FindPersonById).Methods("Get")
	r.HandleFunc("/api/person", controllers.CreatePerson).Methods("Post")
	r.HandleFunc("/api/person/{id}", controllers.DeletePerson).Methods("Delete")
	r.HandleFunc("/api/person/{id}", controllers.EditPerson).Methods("Put")

	log.Fatal(http.ListenAndServe(":8080", r))
}
