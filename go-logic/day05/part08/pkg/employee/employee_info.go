package employee

import (
	"encoding/json"
	"fmt"
)

func (e *Employee) ToString() string {
	return fmt.Sprintf("FullName : %s %s, HireDate : %s, Salary : %2.f",
		e.firstName, e.lastName, e.hireDate.Format("2026-01-26"), e.salary)
}

func (e *Employee) ToJson() (string, error) {
	data := map[string]any{
		"id" 		: e.id,
		"firstName" : e.firstName,
		"lastName" 	: e.lastName,
		"hireData" 	: e.hireDate.Format("2025-12-01"),
		"salary" 	: e.salary,
	}
	jsonBytes, err := json.MarshalIndent(data, "", " ")
	return string(jsonBytes), err
}