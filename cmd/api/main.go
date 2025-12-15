package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BrunnoQ/cadastro-pessoas/internal/application/usecases"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/config"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/logger"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/persistence/mongodb"
	httpHandler "github.com/BrunnoQ/cadastro-pessoas/internal/presentation/http"
	"github.com/BrunnoQ/cadastro-pessoas/internal/presentation/http/handlers"
	"github.com/BrunnoQ/cadastro-pessoas/internal/presentation/http/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "local"
	}

	cfg, err := config.Load(environment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log, err := logger.New(logger.Config{
		Level:  cfg.Logging.Level,
		Format: cfg.Logging.Format,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	log.Info("Starting application",
		zap.String("environment", environment),
		zap.String("version", cfg.App.Version),
	)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	dbConn, err := mongodb.NewConnection(ctx, mongodb.ConnectionConfig{
		URI:             cfg.MongoDB.URI,
		Database:        cfg.MongoDB.Database,
		Timeout:         cfg.MongoDB.Timeout,
		MinPoolSize:     cfg.MongoDB.MinPoolSize,
		MaxPoolSize:     cfg.MongoDB.MaxPoolSize,
		MaxConnIdleTime: cfg.MongoDB.MaxConnIdleTime,
	})
	if err != nil {
		log.Fatal("Failed to connect to MongoDB", zap.Error(err))
	}
	defer dbConn.Close(context.Background())

	log.Info("Connected to MongoDB",
		zap.String("database", cfg.MongoDB.Database),
	)

	// Initialize database indexes
	if err := mongodb.InitializeIndexes(ctx, dbConn.Database); err != nil {
		log.Fatal("Failed to initialize database indexes", zap.Error(err))
	}

	log.Info("Database indexes initialized")

	// Initialize repositories
	personRepo := mongodb.NewMongoPersonRepository(dbConn.Database)

	// Initialize use cases
	createPersonUseCase := usecases.NewCreatePersonUseCase(personRepo, log)
	getPersonUseCase := usecases.NewGetPersonUseCase(personRepo, log)
	listPersonsUseCase := usecases.NewListPersonsUseCase(personRepo, log)
	updatePersonUseCase := usecases.NewUpdatePersonUseCase(personRepo, log)
	deletePersonUseCase := usecases.NewDeletePersonUseCase(personRepo, log)

	// Set Gin mode
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(dbConn)
	personHandler := handlers.NewPersonHandler(createPersonUseCase, getPersonUseCase, listPersonsUseCase, updatePersonUseCase, deletePersonUseCase, log)

	// Create HTTP router
	router := httpHandler.NewRouter(healthHandler, personHandler)
	engine := router.Engine()

	// Apply middleware (order matters!)
	engine.Use(middleware.RecoveryMiddleware(log))
	engine.Use(middleware.RequestIDMiddleware())
	engine.Use(middleware.LoggingMiddleware(log))
	engine.Use(middleware.SecurityHeadersMiddleware())
	engine.Use(middleware.CompressionMiddleware())
	engine.Use(middleware.CORSMiddleware())
	engine.Use(middleware.ValidationMiddleware())

	// Setup routes
	router.SetupRoutes()

	// Create HTTP server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:        engine,
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		IdleTimeout:    cfg.Server.IdleTimeout,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	// Start server in goroutine
	go func() {
		log.Info("HTTP server starting",
			zap.Int("port", cfg.Server.Port),
			zap.String("mode", cfg.Server.Mode),
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start HTTP server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", zap.Error(err))
	}

	log.Info("Server exited")
}
