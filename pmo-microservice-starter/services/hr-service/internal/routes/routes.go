package routes

import (
	"pmo/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// internal/routes/routes.go
func SetupRoutes(db *gorm.DB, validate *validator.Validate) *gin.Engine {
	router := gin.Default()
	//global middleware
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.CORS())
	//set api group
	basePath := viper.GetString("SERVER.BASE_PATH")
	if basePath == "" {
		basePath = "/api"
	}
	// Initialize storage versi hardcode
	/* storage := storage.NewStorageService(
	   "../../storage/uploads", // disimpan di root pmo
	   microservices/storage/uploads
	   "/uploads", // Base URL for accessing files
	   ) */
	// Initialize storage
	// storage := storage.NewStorageService(
	// 	viper.GetString("STORAGE.STORAGE_PATH"), // Base path for file storage,
	// 	viper.GetString("STORAGE.STORAGE_URL"),  // Base URL for accessing files
	// )
	// Serve static files, supaya bisa dical url photonya
	//router.Static("/uploads", "../../storage/uploads") --versi hardcode
	router.Static(viper.GetString("STORAGE.STORAGE_URL"),
		viper.GetString("STORAGE.STORAGE_PATH"))
	api := router.Group(basePath)
	{
		// Deparatments routes
		RegisterDepartmentRoutes(api, db, validate)
		// Buka comment jika feature employee udah dicreate Employee routes
		//RegisterEmployeeRoutes(api, db, validate, storage)
	}
	// Health check
	router.GET("/hr/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy restapi",
		})
	})
	return router
}
