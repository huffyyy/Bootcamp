package routes

import (
	createdepartment "pmo/services/hr-service/features/department/create_department"
	deletedapartment "pmo/services/hr-service/features/department/delete_department"
	getalldepartment "pmo/services/hr-service/features/department/get_all_department"
	getdepartmentbyid "pmo/services/hr-service/features/department/get_department_by_id"
	getdepartmentbyname "pmo/services/hr-service/features/department/get_department_by_name"
	updatedepartment "pmo/services/hr-service/features/department/update_department"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func RegisterDepartmentRoutes(router *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	// Register all department feature routes
	createdepartment.RegisterRoutes(router, db, validate)
	getalldepartment.RegisterRoutes(router, db, validate)
	getdepartmentbyid.RegisterRoutes(router, db, validate)
	getdepartmentbyname.RegisterRoutes(router, db, validate)
	updatedepartment.RegisterRoutes(router, db, validate)
	deletedapartment.RegisterRoutes(router, db, validate)

}
