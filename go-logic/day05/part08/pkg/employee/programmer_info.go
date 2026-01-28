package employee

import (
	"encoding/json"
	"fmt"
)

func (p *Programmer) ToString() string {
	return fmt.Sprintf("Id : %d, FullName : %s %s, HireDate : %s, Salary : %2.f, Placement: %s",
		p.id, p.firstName, p.lastName, p.hireDate.Format("2026-01-26"), p.salary, p.Placement)
}

func (p *Programmer) ToJson() (string, error) {
	data := map[string]any{
		"id" 			: p.id,
		"firstName" 	: p.firstName,
		"lastName" 		: p.lastName,
		"hireData" 		: p.hireDate.Format("2025-12-01"),
		"salary" 		: p.salary,
		"placement" 	: p.Placement,
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")
	return string(jsonBytes), err
}