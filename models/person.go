package models

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Document string `json:"document"`
}

var People []Person
