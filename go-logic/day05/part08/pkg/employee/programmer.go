package employee

import (
	"time"

	"codeid.day05.part08/pkg/departement"
)

type Programmer struct {
	Employee
	Placement Placement
}

func NewProgrammer(firstName string, lastName string, hireDate time.Time, salary float64, departement *departement.Departement, placement Placement) *Programmer {
	return &Programmer{
		Employee: *NewEmployeeWithDept(firstName, lastName, hireDate, salary, departement),
		Placement: placement,
	}
}

func (p *Programmer) GetPlacement() Placement {
	if p != nil {
		return p.Placement
	}
	return ""
}

func (p *Programmer) setPlacement (placement Placement)  {
	if p != nil {
		p.Placement = placement
	}
}

