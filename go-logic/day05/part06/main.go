package main

import (
	"fmt"
	"time"

	"codeid.day05.part06/pkg/departement"
	"codeid.day05.part06/pkg/employee"
)

func main() {
	
	// 1. initial value departement
	deptIt := departement.NewDepartement(1, "IT")
	deptFinance := departement.NewDepartement(2, "Finance")
	deptSales := departement.NewDepartement(3, "Sales")

	// 1. constructor return pointer employee (recomended)
	emp1 := employee.NewEmployeeWithDept("Husnul", "Fikri", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptIt) 
	emp2 := employee.NewEmployeeWithDept("Abdul", "Kareem", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptIt) 
	emp3 := employee.NewEmployeeWithDept("Rini", "Maharani", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptFinance) 
	emp4 := employee.NewEmployeeWithDept("Rasasi", "Hawas", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptFinance)
	emp5 := employee.NewEmployeeWithDept("Naila", "Salsabila", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptSales)

	employees := []*employee.Employee{emp1, emp2, emp3, emp4, emp5}

	for i, v := range employees {
		fmt.Printf("Emp[%d] Addres[%p] Value[%v]\n", i, v, v)
	}
}