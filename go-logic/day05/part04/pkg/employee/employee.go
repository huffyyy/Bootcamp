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
// encapsulation method
func NewEmployee(firstName string, lastName string, hireDate time.Time, salary float64) *Employee  {
	return &Employee{
		firstName: firstName,
		lastName: lastName,
		hireDate: hireDate,
		salary: salary,
	}
}

// 2. constructor return value employee
func NewEmployeeValue(firstName string, lastName string, hireDate time.Time, salary float64) Employee  {
	return Employee{
		firstName: firstName,
		lastName: lastName,
		hireDate: hireDate,
		salary: salary,
	}
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