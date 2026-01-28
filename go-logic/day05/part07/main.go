package main

import (
	"fmt"
	"time"

	"codeid.day05.part07/pkg/departement"
	"codeid.day05.part07/pkg/employee"
)

func main() {
	
	// 1. initial value departement
	deptIt := departement.NewDepartement(1, "IT")
	deptFinance := departement.NewDepartement(2, "Finance")
	deptSales := departement.NewDepartement(3, "Sales")

	// 1. constructor return pointer employee (recomended)
	emp1 := employee.NewEmployeeWithDept("Husnul", "Fikri", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptIt) 
	emp2 := employee.NewProgrammer("Abdul", "Kareem", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptFinance, "BCA") 
	emp3 := employee.NewProgrammer("Rini", "Maharani", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptIt, "IT") 
	emp4 := employee.NewProgrammer("Rasasi", "Hawas", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptFinance, "BRI") 
	emp5 := employee.NewManager("Naila", "Salsabila", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000, deptSales, 10)

	// 2. toString() hanya bisa menampilkan data employee, informasi tambahan seperti placement & totalStaff tidak muncul
	employees := []*employee.Employee{emp1, &emp2.Employee, &emp3.Employee, &emp4.Employee, &emp5.Employee}

	for _, emp := range employees {
		fmt.Println(emp.ToString())
	}

	// 3. populate with interface
    employeesInf := []employee.Info{emp1, emp2, emp3, emp4, emp5}
    for _, v := range employeesInf {
        fmt.Println(v.ToString())
    }

/* 	fmt.Println(emp1.ToString())
	fmt.Println(emp2.ToString())
	fmt.Println(emp3.ToString())
	fmt.Println(emp4.ToString())
	fmt.Println(emp5.ToString()) */
}