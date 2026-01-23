package models

type Department struct {
	departmentId   int
	departmentName string
}

func NewDepartment(departmentId int, departmentName string) *Department {
	return &Department{
		departmentId,
		departmentName,
	}
}