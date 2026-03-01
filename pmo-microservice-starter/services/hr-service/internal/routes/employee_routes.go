package routes

import (
	"pmo/internal/pkg/storage"
	createemployee "pmo/services/hr-service/features/employee/create_employee"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterEmployeeRoutes(router *gin.RouterGroup, db *gorm.DB, validate *validator.Validate,
	storage *storage.StorageService) {
	// Register all employee feature routes
	createemployee.RegisterRoutes(router, db, validate, storage)
}
