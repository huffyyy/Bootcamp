package dto

import "pmo/services/hr-service/internal/domain/models"

type DepartmentResponse struct {
	DepartmentID   int32  `json:"department_id"`
	DepartmentName string `json:"department_name"`
	LocationID     *int32 `json:"location_id,omitempty"`
}

func ToDepartmentResponse(dept *models.Department) *DepartmentResponse {
	if dept == nil {
		return nil
	}
	return &DepartmentResponse{
		DepartmentID:   dept.DepartmentID,
		DepartmentName: dept.DepartmentName,
		LocationID:     dept.LocationID,
	}
}
func ToDepartmentResponses(depts []models.Department) []DepartmentResponse {
	responses := make([]DepartmentResponse, len(depts))
	for i, dept := range depts {
		responses[i] = *ToDepartmentResponse(&dept)
	}
	return responses
}

type DepartmentListResponse struct {
	Departments []DepartmentResponse `json:"departments"`
	Total       int64                `json:"total"`
}

func ToDepartmentListResponse(depts []models.Department, total int64) *DepartmentListResponse {
	return &DepartmentListResponse{
		Departments: ToDepartmentResponses(depts),
		Total:       total,
	}
}
