package main

import (
	"fmt"
	"time"
)

// 1. declare employee
type employee struct {
	employee int
	firstName string
	lastName string
	hireDate time.Time
	salary float64
}

func infoEmployee(e employee) string  {
	return fmt.Sprintf("FullName : %s, %s, HireDate : %s, Salary : %.2f",
		e.firstName, e.lastName, e.hireDate.Format("2026-02-02"), e.salary)
}

func main() {
	// 1. create 2 object struct employee
	emp1 := employee{
		firstName: "Steven", 
		lastName: "King", 
		salary: 24000,
		hireDate: time.Now()}
	emp2 := employee{
		firstName: "Neena",
		lastName: "Kochhar",
		salary: 17000, 
		hireDate: time.Now()}

	emp3 := new(employee)
	emp3.firstName = "Lex"
	emp3.lastName = "De Haan"
	emp3.salary = 1700
	emp3.hireDate = time.Now() 

	// 2. buat slices employees
	employees := []employee{emp1, emp2, *emp3}

	// 3. tampilkan 
	for _, v := range employees {
		fmt.Println(infoEmployee(v))
	}
}