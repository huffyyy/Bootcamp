package routes

import (
	createdepartment "pmo/services/hr-service/features/department/shared/create_department"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterDepartmentRoutes(router *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	// Register all department feature routes
	createdepartment.RegisterRoutes(router, db, validate)
}
