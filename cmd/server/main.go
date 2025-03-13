package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Okabe-Junya/api-template-go/internal/db"
	apierror "github.com/Okabe-Junya/api-template-go/internal/error"
	"github.com/Okabe-Junya/api-template-go/internal/handlers"
	"github.com/Okabe-Junya/api-template-go/internal/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

func main() {
	// Initialize logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Set environment variables (in production, these should be properly configured)
	dbConnectionString := os.Getenv("DB_CONNECTION")
	if dbConnectionString == "" {
		dbConnectionString = "user:password@tcp(localhost:3306)/db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize database connection
	conn, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// Test connection before setting up defer
	if err := conn.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Database connection test failed")
	}

	// Set up defer after successful connection
	defer conn.Close()

	log.Info().Msg("Successfully connected to database")

	// Initialize database queries
	query := db.New(conn)

	// Initialize Gin (use ReleaseMode in production)
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.New()

	// Set up middleware
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(apierror.ErrorHandlingMiddleware())
	router.Use(middleware.PrometheusMetricsMiddleware())
	router.Use(gin.Recovery())

	// Register Prometheus metrics endpoint
	middleware.RegisterPrometheusHandler(router)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// Initialize handlers
	userHandler := handlers.NewUserHandler(query)
	itemHandler := handlers.NewItemHandler(query)
	userItemHandler := handlers.NewUserItemHandler(query)

	// API v1 route group
	v1 := router.Group("/api/v1")
	{
		// Sample route
		v1.GET("/", handlers.SampleHandler)

		// User routes
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// Item routes
		items := v1.Group("/items")
		{
			items.POST("", itemHandler.CreateItem)
			items.GET("/:id", itemHandler.GetItem)
			items.DELETE("/:id", itemHandler.DeleteItem)
		}

		// User-Item routes
		userItems := v1.Group("/user_items")
		{
			userItems.POST("", userItemHandler.CreateUserItem)
			userItems.GET("/:user_id/:item_id", userItemHandler.GetUserItem)
			userItems.DELETE("/:user_id/:item_id", userItemHandler.DeleteUserItem)
		}
	}

	// Legacy routes for backward compatibility
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	router.POST("/items", itemHandler.CreateItem)
	router.GET("/items/:id", itemHandler.GetItem)
	router.DELETE("/items/:id", itemHandler.DeleteItem)
	router.POST("/user_items", userItemHandler.CreateUserItem)
	router.GET("/user_items/:user_id/:item_id", userItemHandler.GetUserItem)
	router.DELETE("/user_items/:user_id/:item_id", userItemHandler.DeleteUserItem)

	// Initialize server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// Set up graceful shutdown
	go func() {
		log.Info().Msgf("Server is running on port: %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Initiating shutdown...")

	// Shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server shutdown failed")
		// Let defer run before returning from main
		return
	}

	log.Info().Msg("Server shutdown completed")
}
