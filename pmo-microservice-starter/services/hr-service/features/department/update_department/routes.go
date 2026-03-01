package updatedepartment

import (
	"pmo/services/hr-service/features/department/shared/repository"
	"pmo/services/hr-service/features/department/shared/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB, v *validator.Validate) {
	repo := repository.NewDepartmentRepository(db)
	deptValidator := validators.NewDepartmentValidator(v)
	handler := NewUpdateDepartmentHandler(repo, deptValidator)
	router.PUT("/departments/:id", handler.Handle)
}
