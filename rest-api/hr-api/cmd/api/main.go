package main

import (
	"log"
	"net/http"

	"github.com/codeid/hr-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. set datasourcename db config
	_, err := database.SetupDB()
	if err != nil {
		log.Fatal("failed to connect db %w", err)
	}

	// 3. setup router
	router := gin.Default()

	// 3.1 call handler
	router.GET("/", hellowordHandler)

	// 4. run webserver
	router.Run(":8080")
}

func hellowordHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H {
		"message" : "Hello World",
		"status" : "running",
	})
}