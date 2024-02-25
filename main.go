package main

import (
	"github.com/br-jeff/go-api-http/models"
	"github.com/br-jeff/go-api-http/routes"
)

func main() {
	models.People = []models.Person{
		{Id: 1, Name: "Neymar Santos", Document: "12345435"},
		{Id: 2, Name: "Lionel Messi", Document: "2134234"},
	}

	routes.HandleRequest()
}
