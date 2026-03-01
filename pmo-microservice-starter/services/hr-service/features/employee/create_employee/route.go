package createemployee

import (
	"pmo/internal/pkg/storage"
	"pmo/services/hr-service/features/employee/shared/repository"
	"pmo/services/hr-service/features/employee/shared/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB, validate *validator.Validate, storage *storage.StorageService,
) {
	repo := repository.NewEmployeeRepository(db)
	empValidator := validators.NewEmployeeValidator(validate)
	handler := NewCreateEmployeeHandler(repo, empValidator, storage)
	router.POST("/employees", handler.Handle)
}
