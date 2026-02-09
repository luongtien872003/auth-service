package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/luongtien872003/login/internal/handler"
	"github.com/luongtien872003/login/internal/repository"
)

func main() {
	// 1. Connect DB
	db, err := sql.Open(
		"postgres",
		"postgres://auth:secret@localhost:5432/authdb?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// 2. Init Repository
	userRepo := repository.NewUserRepository(db)

	// 3. Init Handler
	authHandler := handler.NewAuthHandler(userRepo)

	// 4. Init Router
	r := gin.Default()

	// 5. Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 6. Auth routes
	// POST /register - Register a new user
	r.POST("/register", authHandler.Register)

	// 7. Start server
	log.Println("Server starting on :8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
