package models

import (
	"fmt"
	"time"
)

// 1. declare employee
// employee klo diawalin huruf kecil : private
// employee prefix huruf besar : publi
// implementasi encapsulation
type Employee struct {
	employeeid int
	firstName string
	lastName string
	hireDate time.Time
	salary float64
	Departement Department
}

// constructor : initial value employee
// constructor : selalu return pointer employee
func NewEmployee(employeeid int, firstName string, lastName string, hireDate time.Time, salary float64, departement Department) *Employee  {
	// return address
	return &Employee{
		employeeid,
		firstName,
		lastName,
		hireDate,
		salary,
		departement,
	}
 }

func NewEmployee2(employeeid int, firstName string, lastName string, hireDate time.Time) *Employee {
	return &Employee{
		employeeid: employeeid,
		firstName:  firstName,
		lastName:   lastName,
		hireDate:   hireDate,
	}
}

// method reciver
func (e *Employee) Salary(salary float64) error {
	if  salary > 10_000 {
		return  fmt.Errorf("%.2f is to high", salary)
	}
	e.salary = salary
	return nil
}

func (e *Employee) InfoEmployee() string {
	return fmt.Sprintf("Employee :[%s %s %.2f %s]", e.firstName, e.lastName, e.salary, e.Departement.departmentName)
}

