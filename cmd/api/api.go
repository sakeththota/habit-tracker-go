package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sakeththota/habit-tracker-go/service/habit"
	"github.com/sakeththota/habit-tracker-go/service/progress"
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
	v2 := router.Group("/v2")

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(v2)

	habitStore := habit.NewStore(s.db)
	habitHandler := habit.NewHandler(habitStore, userStore)
	habitHandler.RegisterRoutes(v2)

	progressStore := progress.NewStore(s.db)
	progressHandler := progress.NewHandler(progressStore, userStore, habitStore)
	progressHandler.RegisterRoutes(v2)

	log.Println("Listening on", s.addr)
	return router.Run(s.addr)
}
