package main

import (
	"log"

	"github.com/sakeththota/habit-tracker-go/cmd/api"
	"github.com/sakeththota/habit-tracker-go/config"
	"github.com/sakeththota/habit-tracker-go/db"
)

func main() {
	dbpool, err := db.NewPgxPool(db.PgxConfig{
		Database: config.Envs.DBName,
		Password: config.Envs.DBPassword,
		Username: config.Envs.DBUsername,
		Port:     config.Envs.DBPort,
		Host:     config.Envs.DBHost,
		Schema:   config.Envs.DBSchema,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	server := api.NewApiServer(":8080", dbpool)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
