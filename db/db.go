package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxConfig struct {
	Name     string
	Password string
	Username string
	Port     string
	Host     string
	Schema   string
}

func NewPostgreSQLStorage(cfg PgxConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", cfg.Name, cfg.Password, cfg.Username, cfg.Port, cfg.Host, cfg.Schema)

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer dbpool.Close()

	return dbpool, nil
}
