package getalldepartment

import "pmo/services/hr-service/features/department/shared/repository"

type GetAllDepartmentsQuery struct {
	Page           int
	PageSize       int
	SortBy         string
	SortDir        string
	DepartmentName string
	LocationID     *int32
}

func NewGetAllDepartmentsQuery(page, pageSize int, sortBy, sortDir string,
	deptName string, locationID *int32) GetAllDepartmentsQuery {
	return GetAllDepartmentsQuery{
		Page:           page,
		PageSize:       pageSize,
		SortBy:         sortBy,
		SortDir:        sortDir,
		DepartmentName: deptName,
		LocationID:     locationID,
	}
}
func (q GetAllDepartmentsQuery) ToRepositoryParams() *repository.FindAllParams {
	return &repository.FindAllParams{
		Page:           q.Page,
		PageSize:       q.PageSize,
		SortBy:         q.SortBy,
		SortDir:        q.SortDir,
		DepartmentName: q.DepartmentName,
		LocationID:     q.LocationID,
	}
}
