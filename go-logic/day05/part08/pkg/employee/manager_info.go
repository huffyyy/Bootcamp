package employee

import (
	"encoding/json"
	"fmt"
)

func (m *Manager) ToString() string {
	return fmt.Sprintf("Id : %d, FullName : %s %s, HireDate : %s, Salary : %2.f, TotalStaff: %d",
		m.id, m.firstName, m.lastName, m.hireDate.Format("2026-01-26"), m.salary, m.totalStaff)
}

func (m *Manager) ToJson() (string, error) {
	data := map[string]any{
		"id" 			: m.id,
		"firstName" 	: m.firstName,
		"lastName" 		: m.lastName,
		"hireData" 		: m.hireDate.Format("2025-12-01"),
		"salary" 		: m.salary,
		"totalStaff" 	: m.totalStaff,
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")
	return string(jsonBytes), err
}