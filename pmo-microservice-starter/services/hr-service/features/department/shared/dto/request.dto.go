package dto

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name" validate:"required,min=2,max=30"`
	LocationID     *int32 `json:"location_id" validate:"omitempty,min=1"`
}

type UpdateDepartmentRequest struct {
	DepartmentName string `json:"department_name" validate:"omitempty,min=2,max=30"`
	LocationID     *int32 `json:"location_id" validate:"omitempty,min=1"`
}

type GetDepartmentByIDRequest struct {
	DepartmentID int32 `uri:"id" validate:"required,min=1"`
}

type GetDepartmentByNameRequest struct {
	DepartmentName string `form:"name" validate:"required,min=2,max=30"`
}

type DeleteDepartmentRequest struct {
	DepartmentID int32 `uri:"id" validate:"required,min=1"`
}
type GetAllDepartmentsRequest struct {
	Page     int    `form:"page" validate:"omitempty,min=1"`
	PageSize int    `form:"page_size" validate:"omitempty,min=1,max=100"`
	SortBy   string `form:"sort_by" validate:"omitempty,oneof=department_id department_name location_id"`
	SortDir  string `form:"sort_dir" validate:"omitempty,oneof=asc desc"`

	// Filters
	DepartmentName string `form:"department_name" validate:"omitempty,min=1,max=30"`
	LocationID     *int32 `form:"location_id" validate:"omitempty,min=1"`
}

func (r *GetAllDepartmentsRequest) SetDefaults() {
	if r.Page <= 0 {
		r.Page = 1
	}
	if r.PageSize <= 0 {
		r.PageSize = 10
	}
	if r.SortBy == "" {
		r.SortBy = "department_id"
	}
	if r.SortDir == "" {
		r.SortDir = "asc"
	}
}
