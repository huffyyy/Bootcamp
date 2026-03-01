package getdepartmentbyid

type GetDepartmentByIDQuery struct {
	DepartmentID int32
}

func NewGetDepartmentByIDQuery(id int32) GetDepartmentByIDQuery {
	return GetDepartmentByIDQuery{
		DepartmentID: id,
	}
}
