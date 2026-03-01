package getdepartmentbyname

type GetDepartmentByNameQuery struct {
	DepartmentName string
}

func NewGetDepartmentByNameQuery(name string) GetDepartmentByNameQuery {
	return GetDepartmentByNameQuery{
		DepartmentName: name,
	}
}
