package employee

import "time"

// 1. jadikan employee public
type Employee struct {
	id        int64
	firstName string
	lastName  string
	hireDate  time.Time
	salary float64
}

// 1. constructor return pointer employee
// sharing value
func NewEmployee(firstName string, lastName string, hireDate time.Time, salary float64) *Employee  {
	return &Employee{
		firstName: firstName,
		lastName: lastName,
		hireDate: hireDate,
		salary: salary,
	}
}

