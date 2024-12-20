package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sakeththota/habit-tracker-go/config"
	"github.com/sakeththota/habit-tracker-go/db"
)

func main() {
	db, err := db.NewPgxDb(db.PgxConfig{
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

	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		log.Fatalf("Unable to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "pgx", driver)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
