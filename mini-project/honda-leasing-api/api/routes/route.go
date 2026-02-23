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
	ProductRepo := repositories.NewLeasingProductRepository(db)
	customerRepo := repositories.NewCustomerRepository(db)
	contractRepo := repositories.NewLeasingContractRepository(db)
	taskRepo := repositories.NewLeasingTaskRepository(db)
	taskAtrRepo := repositories.NewLeasingTaskAtrRepository(db)

	//6. Initialize services
	motorService := services.NewMotorService(motorRepo)
	simulationService := services.NewSimulationService(motorRepo, ProductRepo)
	draftService := services.NewOrderDraftService(motorRepo, ProductRepo, customerRepo, contractRepo)
	progressService := services.NewOrderProgressService(contractRepo, taskRepo)
	leasingService := services.NewLeasingService(contractRepo, customerRepo, motorRepo, ProductRepo, taskRepo, taskAtrRepo)

	//7. Initialize handlers
	motorHandler := handlers.NewMotorHandler(motorService)
	simulationHandler := handlers.NewSimulationHandler(simulationService)
	draftHandler := handlers.NewOrderDraftHandler(draftService)
	progressHandler := handlers.NewOrderProgressHandler(progressService)
	leasingHandler := handlers.NewLeasingHandler(leasingService)

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

		simulations := api.Group("/simulations")
		{
			simulations.POST("", simulationHandler.Simulate)
		}

		orders := api.Group("/orders")
		{
			orders.POST("/draft", draftHandler.CreateOrderDraft)
			orders.GET("/:id", progressHandler.GetOrderProgress)
			orders.PATCH("/task/:task_id", progressHandler.UpdateTaskStatus)
		}

		leasings := api.Group("/leasing")
		{
			leasings.GET("/inbox", leasingHandler.GetInboxByStatus)
			leasings.GET("/contracts/:id", leasingHandler.GetRequestDetail)
			leasings.GET("/:id/progress-detail", leasingHandler.GetProgressDetail)
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
