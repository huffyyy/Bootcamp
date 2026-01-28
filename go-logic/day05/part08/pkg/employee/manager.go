package employee

import (
	"time"

	"codeid.day05.part08/pkg/departement"
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

