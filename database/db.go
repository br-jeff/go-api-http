package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDatabase() {
	connectionString := "host=localhost user=dev password=dev dbname=dev port=5432 sslmode=disable "

	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Erro when try to connect into database")
	}
}
