package employee

import (
	"fmt"
	"time"

	"codeid.day05.part07/pkg/departement"
)

type Programmer struct {
	Employee
	placement string
}

func NewProgrammer(firstName string, lastName string, hireDate time.Time, salary float64, departement *departement.Departement, placement string) *Programmer {
	return &Programmer{
		Employee: *NewEmployeeWithDept(firstName, lastName, hireDate, salary, departement),
		placement: placement,
	}
}

func (p *Programmer) GetPlacement() string {
	if p != nil {
		return p.placement
	}
	return ""
}

func (p *Programmer) setPlacement (placement string)  {
	if p != nil {
		p.placement = placement
	}
}

func (p *Programmer) ToString() string {
	return fmt.Sprintf("Id : %d, FullName : %s %s, HireDate : %s, Salary : %2.f, Placement: %s",
			p.id ,p.firstName, p.lastName, p.hireDate.Format("2026-01-26"), p.salary, p.placement )
}

func (p *Programmer) ToJson() (string, error)  {
	panic("not implemented")
}