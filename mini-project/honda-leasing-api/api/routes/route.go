package routes

import (
	"net/http"

	"github.com/codeid/honda-leasing-api/internal/handlers"
	"github.com/codeid/honda-leasing-api/internal/repositories"
	"github.com/codeid/honda-leasing-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	//5. Initialize repositories
	motorRepo := repositories.NewMotorRepository(db)

	//6. Initialize services
	motorService := services.NewMotorService(motorRepo)

	//7. Initialize handlers
	motorHandler := handlers.NewMotorHandler(motorService)

	//3.1 call handler
	router.GET("/", welcomeHandler)

	//9.call basepath
	basePath := viper.GetString("SERVER.BASE_PATH")
	if basePath == "" {
		basePath = "/api/v1"
	}

	api := router.Group(basePath)
	{
		motors := api.Group("/motors")
		{
			motors.GET("", motorHandler.GetAllMotor)
			motors.GET("/search", motorHandler.SearchMotors)
			motors.GET("/category", motorHandler.GetMotorByCategory)
			motors.GET("/:id", motorHandler.GetMotorByID)
		}
	}

}

// 2. create first handler
func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to gin framework",
		"status":  "running",
	})
}
