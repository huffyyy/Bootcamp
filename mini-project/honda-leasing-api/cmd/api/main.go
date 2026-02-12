package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codeid/honda-leasing-api/api/routes"
	"github.com/codeid/honda-leasing-api/internal/configs"
	"github.com/codeid/honda-leasing-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("APP_ENV", "development")
	config := configs.Load()

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer database.CloseDB(db)

	// OPTIONAL: AutoMigrate (only for dev)
	/* 	if config.Environment == "development" {
		if err := database.InitAutoMigrate(db); err != nil {
			log.Fatalf("auto migrate failed: %v", err)
		}
	} */

	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	routes.SetupRoutes(router, db.DB)

	server := &http.Server{
		Addr:    config.Server.Address,
		Handler: router,
	}

	go func() {
		log.Printf("Server running on %s (%s mode)",
			config.Server.Address,
			config.Environment,
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("Server exited properly")
}
