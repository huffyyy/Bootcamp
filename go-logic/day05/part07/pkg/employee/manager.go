package employee

import (
	"fmt"
	"time"

	"codeid.day05.part07/pkg/departement"
)

type Manager struct {
	Employee
	totalStaff int64
}

func NewManager(firstName string, lastName string, hireDate time.Time, salary float64, departement *departement.Departement, totalStaff int64) *Manager {
	return &Manager{
		Employee: *NewEmployeeWithDept(firstName, lastName, hireDate, salary, departement),
		totalStaff: totalStaff,
	}
}

func (m *Manager) getTotalStaff() int64 {
	if m != nil {
		return m.totalStaff
	}
	return 0
}

func (m *Manager) setTotalStaff(totalStaff int64) {
	if m != nil {
		m.totalStaff = totalStaff
	}
}

func (m *Manager) ToString() string  {
	return fmt.Sprintf("Id : %d, FullName : %s %s, HireDate : %s, Salary : %2.f, TotalStaff: %d",
			m.id ,m.firstName, m.lastName, m.hireDate.Format("2026-01-26"), m.salary, m.totalStaff )
}

func (m *Manager) ToJson() (string, error)  {
	panic("not implemented")
}
