package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxConfig struct {
	Database string
	Password string
	Username string
	Port     string
	Host     string
	Schema   string
}

func NewPgxPool(cfg PgxConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Schema)
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	return dbpool, nil
}

func NewPgxDb(cfg PgxConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Schema)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	return db, nil
}
