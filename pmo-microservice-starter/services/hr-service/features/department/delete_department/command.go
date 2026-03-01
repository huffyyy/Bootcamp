package deletedapartment

type DeleteDepartmentCommand struct {
	DepartmentID int32
}

func NewDeleteDepartmentCommand(id int32) DeleteDepartmentCommand {
	return DeleteDepartmentCommand{
		DepartmentID: id,
	}
}
