package main

import (
	"log"
	"net/http"

	"github.com/codeid/hr-api/internal/handlers"
	"github.com/codeid/hr-api/internal/repositories"
	"github.com/codeid/hr-api/internal/services"
	"github.com/codeid/hr-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	//1. set datasourcename db config
	db, err := database.SetupDB() //posggre
	if err != nil {
		log.Fatal("failed to connect db:%w", err)
	}

	//1.1 initautomigrate --> baru
	database.InitAutoMigrate(db)

	//init repositories
	regionRepo := repositories.NewRegionRepository(db)

	//init service
	regionService := services.NewRegionService(regionRepo)

	//init handler/controler
	regionHandler := handlers.NewRegionHandler(regionService)

	//3. setup route
	router := gin.Default()

	//4. create router endpoint
	api := router.Group("/api")
	{
		//create region route
		regions := api.Group("/regions")
		{
			regions.GET("", regionHandler.GetRegions)
			regions.GET("/:id", regionHandler.GetRegion)
			regions.POST("", regionHandler.CreateRegion)
			regions.PUT("/:id", regionHandler.UpdateRegion)
			regions.DELETE("/:id", regionHandler.DeleteRegion)
			regions.GET("/countries", regionHandler.GetRegionsWithCountry)
			//add new endpoint
			regions.GET("/:id/countries", regionHandler.GetRegionByIdWithCountry)
		}

		//countries

	}

	//3.1 call handler
	router.GET("/", helloWorldHandler)

	//4. run webserver di port 8080
	router.Run(":8080")
}

func helloWorldHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
		"status":  "running",
	})
}