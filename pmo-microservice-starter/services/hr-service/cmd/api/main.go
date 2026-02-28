package main

import (
	"log"
	"os"
	"os/signal"
	"pmo/services/hr-service/internal/configs"
	"pmo/services/hr-service/internal/database"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	//1. set environment (bisa cmd atau system environment)
	os.Setenv("APP_ENV", "development")
	//2.Load configuration
	config := configs.Load()
	//3 Set Gin mode based on environment
	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	//4. current code lebih minimalis
	db, err := database.InitDB(config)
	if err != nil {
		log.Fatal("failed to initialize database:%w", err)
	}
	//4.1. close db sebelum main application exit
	defer database.CloseDB(db)

	//5.Initialize validator (global/singleton)
	//validate := validator.GetValidator()

	//6.Setup routes, up running gin web server
	//router := routes.SetupRoutes(db.DB, validate)

	//6.1.Start server

	log.Printf("Server starting on %s in %s mode", config.Server.Address,
		config.Environment)

	//6.2.goroutine baru
	// go func() {
	// 	if err := router.Run(config.Server.Address); err != nil && err !=
	// 		http.ErrServerClosed {
	// 		log.Fatalf("Failed to start server: %v", err)
	// 	}
	// }()

	// 1. Graceful shutdown : nunggu operasi selesai baru shutdown server
	// 2. Tanpa Graceful Shutdowon : close connection,
	// tapi ada kemungkinan operasi seperti query masih jalan
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
