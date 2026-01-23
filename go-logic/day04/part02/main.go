package main

import (
	"fmt"
	"time"

	"codeid.day04.part01/part02/models"
)

func main() {
	//create object department
	dept := models.NewDepartment(9, "IT Departement")

	//1. create 2 object struct employee
	emp1 := models.NewEmployee(100, "Steven", "Seagel", time.Now(), 24_000, *dept)
	emp2 := models.NewEmployee2(101, "Nina", "Kochar", time.Now())

	err := 	emp2.Salary(100_000)
	if err != nil {
		fmt.Println("Failed to update salary", err)
	}
	emp2.Salary(5000)

	emp3 := models.NewEmployee(100, "Lex", "Luthor", time.Now(), 18_000, *dept)

	// 2. buat slices employees
	employees := []*models.Employee{emp1, emp2}

	employees = append(employees, emp3)

	// 3. tampilkan 
	for _, v := range employees {
		fmt.Println(v.InfoEmployee())
	}
}