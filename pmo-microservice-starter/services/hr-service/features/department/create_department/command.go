package createdepartment

import "pmo/services/hr-service/internal/domain/models"

type CreateDepartmentCommand struct {
	DepartmentName string
	LocationID     *int32
}

func NewCreateDepartmentCommand(name string, locationID *int32) CreateDepartmentCommand {
	return CreateDepartmentCommand{
		DepartmentName: name,
		LocationID:     locationID,
	}
}
func (c CreateDepartmentCommand) ToModel() *models.Department {
	return &models.Department{
		DepartmentName: c.DepartmentName,
		LocationID:     c.LocationID,
	}
}
