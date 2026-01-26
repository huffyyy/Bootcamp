package main

import (
	"fmt"
	"time"
)

type employee struct {
	firstName string
	lastName  string
	hireDate  time.Time
	salary float64
}

func toString(e employee) string  {
	return fmt.Sprintf("FullName : %s %s, HireDate : %s, Salary : %2.f", 
		e.firstName, e.lastName, e.hireDate.Format("2026-01-26"), e.salary)
}

func main() {

	// 1. literal struct
	emp1 := employee{firstName: "Husnul", lastName: "Fikri", hireDate: time.Now(), salary: 100_000}
	// fmt.Printf("emp1:%v\n", emp1)
	fmt.Println(toString(emp1))

	// 2. urutan field
	emp2 := employee{"Budi", "Juna", time.Now(), 150_000}
	// fmt.Printf("emp2:%v\n", emp2)
	fmt.Println(toString(emp2))

	// 3. using constructor
	emp3 := new(employee)
	emp3.firstName = "Charlie"
	emp3.lastName = "Puth"
	emp3.hireDate = time.Now()
	emp3.salary = 125_000
	// fmt.Printf("emp3:%v\n", emp3)
	fmt.Println(toString(*emp3))

}