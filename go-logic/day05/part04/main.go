package main

import (
	"fmt"
	"time"

	"codeid.day05.part04/pkg/employee"
)

func main() {
	
	// 1. constructor return pointer employee 
	emp1 := employee.NewEmployee("Husnul", "Fikri", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000) 
	emp2 := employee.NewEmployee("Abdul", "Kareem", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000) 

	fmt.Println(emp1)
	fmt.Println(emp2)

	// 1, update emp1 by emp3
	emp3 := emp1
	
	if err := emp3.SetSalary(200_000); err != nil {
		fmt.Println("failed to update salary : ", err)
	}

	fmt.Println("emp1 after : ", emp1)

	// 2. constructor return value employee
	emp4 := employee.NewEmployeeValue("Rini", "Maharani", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000) 

	fmt.Println("emp4 before : ", emp4)

	emp5 := emp4
	
	if err := emp5.SetSalary(650_000); err != nil {
		fmt.Println("failed to update salary : ", err)
	}
	
	emp5.SetFirstName("Wawan")
	emp5.SetLastName("Windah")
	
	fmt.Println("emp4 after : ", emp4)
	
}