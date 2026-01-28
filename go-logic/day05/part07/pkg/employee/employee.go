package employee

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"codeid.day05.part07/pkg/departement"
)

var lastID atomic.Int64
func init()  {
	lastID.Store(100)
}

func GenerateID() int64  {
	return lastID.Add(1)
}

// 1. jadikan employee public
type Employee struct {
	id        	int64
	firstName 	string
	lastName  	string
	hireDate  	time.Time
	salary		float64
	departement *departement.Departement
}

// 1. constructor return pointer employee
// sharing value
// encapsulation method
func NewEmployee(firstName string, lastName string, hireDate time.Time, salary float64) *Employee  {
	return &Employee{
		id: 		GenerateID(),
		firstName: 	firstName,
		lastName: 	lastName,
		hireDate: 	hireDate,
		salary: 	salary,
	}
}

// 2. constructor return value employee
func NewEmployeeValue(firstName string, lastName string, hireDate time.Time, salary float64) Employee  {
	return Employee{
		id: 		GenerateID(),
		firstName: 	firstName,
		lastName: 	lastName,
		hireDate: 	hireDate,
		salary: 	salary,
	}
}

// 2. constructor with departement
func NewEmployeeWithDept(firstName string, lastName string, hireDate time.Time, salary float64, departement *departement.Departement) *Employee  {
	return &Employee{
		id: 			GenerateID(),
		firstName: 		firstName,
		lastName: 		lastName,
		hireDate: 		hireDate,
		salary: 		salary,
		departement: 	departement,

	}
}

func NewEmployeeValid(firstName string, lastName string, hireDate time.Time, salary float64) (*Employee, error)  {
	
	/* if salary < minimunWage {
		return nil, ErrInvalidSalaryMin
	} else if salary > maximumWage {
		return nil, ErrInvalidSalaryMax
	}

	if strings.TrimSpace(firstName) == "" {
		return nil, ErrEmptyFirstName
	} */

	if err := validateEmployee(firstName, lastName, hireDate, salary); err != nil {
		return nil, err
	}
	
	return &Employee{
		id: 		GenerateID(),	
		firstName: 	firstName,
		lastName: 	lastName,
		hireDate: 	hireDate,
		salary: 	salary,
	}, nil
}

func (e *Employee) GetId() int64 {
	if e != nil {
		return e.id
	}
	return 0
}

func (e *Employee) SetId(id int64) {
	if e != nil {
		e.id = id
	}
}

func (e *Employee) GetFirstName() string {
	if e != nil {
		return e.firstName
	}
	return ""
}

func (e *Employee) SetFirstName(firstName string) {
	if e != nil {
		e.firstName = firstName
	}
}

func (e *Employee) GetLastName() string {
	if e != nil {
		return e.lastName
	}
	return ""
}

func (e *Employee) SetLastName(lastName string) {
	if e != nil {
		e.lastName = lastName
	}
}

func (e *Employee) GetHireDate() time.Time {
	var ret time.Time
	if e != nil {
		ret = e.hireDate
	}
	return ret
}

func (e *Employee) SetHireDate(hireDate time.Time) {
	if e != nil {
		e.hireDate = hireDate
	}
}

func (e *Employee) GetSalary() float64 {
	if e != nil {
		return e.salary
	}
	return 0
}

func (e *Employee) SetSalary(salary float64) error {
	if e != nil {
		if salary < minimunWage {
			return ErrInvalidSalaryMin
		} else if salary > maximumWage {
			return ErrInvalidSalaryMax
		}
		e.salary = salary
	}
	return nil
}

func validateEmployee(firstName, lastName string, hireDate time.Time, salary float64) error  {
	if strings.TrimSpace(firstName) == "" {
		return ErrEmptyFirstName
	}

	if strings.TrimSpace(lastName) == "" {
		return ErrEmptyLastName
	}

	if salary < minimunWage {
		return ErrInvalidSalaryMin
	} else if salary > maximumWage {
		return ErrInvalidSalaryMax
	}

	if hireDate.After(time.Now()) {
		return ErrFutureHireDate
	}
	return nil
}

func (e *Employee) ToString() string {
		return fmt.Sprintf("FullName : %s %s, HireDate : %s, Salary : %2.f",
		e.firstName, e.lastName, e.hireDate.Format("2026-01-26"), e.salary)

}

func (e *Employee) ToJson() (string, error) {
	panic("not implemented")

}