package main

import (
	"log"

	"github.com/sakeththota/habit-tracker-go/config"
	"github.com/sakeththota/habit-tracker-go/db"
)

func main() {
	pgxPool, err := db.NewPostgreSQLStorage(db.PgxConfig{
		Name:     config.Envs.DBName,
		Password: config.Envs.DBPassword,
		Username: config.Envs.DBUsername,
		Port:     config.Envs.DBPort,
		Host:     config.Envs.DBHost,
		Schema:   config.Envs.DBSchema,
	})
	if err != nil {
		log.Fatal(err)
	}
}
