package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kalyanbasso/taskhub/internal/config"
	"github.com/kalyanbasso/taskhub/internal/controller"
	"github.com/kalyanbasso/taskhub/internal/repository"
	"github.com/kalyanbasso/taskhub/internal/usecase"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	taskRepo := repository.NewTaskRepository(db)

	// Initialize use cases
	taskUseCase := usecase.NewTaskUseCase(taskRepo)

	// Initialize controllers
	taskController := controller.NewTaskController(taskUseCase)

	// Initialize Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "TaskHub API is running",
		})
	})

	// Register routes
	taskController.RegisterRoutes(r)

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
