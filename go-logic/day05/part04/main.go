package main

import (
	"fmt"
	"time"

	"codeid.day05.part04/pkg/employee"
)

func main() {
	emp1 := employee.NewEmployee("Husnul", "Fikri", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000) 
	emp2 := employee.NewEmployee("Abdul", "Kareem", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000) 

	fmt.Println(emp1)
	fmt.Println(emp2)

	emp3 := emp1
	
	if err := emp3.SetSalary(200_000); err != nil {
		fmt.Println("failed to update salary : ", err)
	}

	fmt.Println("emp1 after : ", emp1)

}