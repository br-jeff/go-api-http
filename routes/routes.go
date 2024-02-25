package routes

import (
	"log"
	"net/http"

	"github.com/br-jeff/go-api-http/controllers"
)

func HandleRequest() {
	// r := mux.NewRouter()
	http.HandleFunc("/", controllers.IndexController)
	http.HandleFunc("/api/person", controllers.ListAllPeople)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
