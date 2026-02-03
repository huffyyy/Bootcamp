package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/codeid/hr-api/api/routes"
	"github.com/codeid/hr-api/internal/configs"
	"github.com/codeid/hr-api/internal/models"
	"github.com/codeid/hr-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//1. set environment (bisa cmd atau system environment)
	os.Setenv("APP_ENV", "development")

	//2.Load configuration
	config := configs.Load()

	//1. current code lebih minimalis
	db, err := database.InitDB(config)
	if err != nil {
		log.Fatal("failed to initialize database:%w", err)
	}
	defer database.CloseDB(db)

	// Run auto migration
	if err := database.AutoMigrate(db, &models.Region{}, &models.Country{}); err != nil {
		log.Printf("Warning: Auto migration failed: %v", err)
	}

	// Set Gin mode based on environment
	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Setup routes
	router := gin.Default()
	routes.SetupRoutes(router, db.DB)

	// Start server
	log.Printf("Server starting on %s in %s mode", config.Server.Address, config.Environment)

	go func() {
	if err := router.Run(config.Server.Address); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 1. Graceful shutdown : nunggu operasi selesai baru shutdown server
	// 2. Tanpa Graceful Shutdowon : close connection,ada kemungkinan operasi seperti query masih jalan
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}