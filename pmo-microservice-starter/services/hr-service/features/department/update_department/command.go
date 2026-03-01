package updatedepartment

import "pmo/services/hr-service/internal/domain/models"

type UpdateDepartmentCommand struct {
	DepartmentID   int32
	DepartmentName string
	LocationID     *int32
}

func NewUpdateDepartmentCommand(id int32, name string, locationID *int32) UpdateDepartmentCommand {
	return UpdateDepartmentCommand{
		DepartmentID:   id,
		DepartmentName: name,
		LocationID:     locationID,
	}
}

func (c UpdateDepartmentCommand) ApplyToModel(dept *models.Department) {
	if c.DepartmentName != "" {
		dept.DepartmentName = c.DepartmentName
	}
	if c.LocationID != nil {
		dept.LocationID = c.LocationID
	}
}
