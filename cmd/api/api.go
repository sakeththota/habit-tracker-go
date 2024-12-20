package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sakeththota/habit-tracker-go/service/habit"
	"github.com/sakeththota/habit-tracker-go/service/user"
)

type APIServer struct {
	addr string
	db   *pgxpool.Pool
}

func NewApiServer(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(v1)

	habitStore := habit.NewStore(s.db)
	habitHandler := habit.NewHandler(habitStore)
	habitHandler.RegisterRoutes(v1)

	log.Println("Listening on", s.addr)
	return router.Run(s.addr)
}
