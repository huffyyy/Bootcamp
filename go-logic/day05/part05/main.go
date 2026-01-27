package main

import (
	"fmt"
	"time"

	"codeid.day05.part05/pkg/employee"
)

func main() {
	
	// 1. constructor return pointer employee (recomended)
	emp1 := employee.NewEmployee("Husnul", "Fikri", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000) 
	emp2 := employee.NewEmployee("Abdul", "Kareem", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000) 
	emp3 := employee.NewEmployee("Rini", "Maharani", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000) 
	emp4 := employee.NewEmployee("Rasasi", "Hawas", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000)
	emp5 := employee.NewEmployee("Naila", "Salsabila", time.Date(2026, time.January, 21, 0, 0, 0, 0, time.UTC), 500_000)

	// 2. append to slice pointer employee (perfomance)
	employeesPtr := []*employee.Employee{emp1, emp2, emp3, emp4, emp5}

	for i, v := range employeesPtr {
		fmt.Printf("Emp[%d] Addres[%p] Value[%v]\n", i, v, v)
	}

	// 3. update value object to employee[0]
	emp1Update := employeesPtr[0]
	emp1Update.SetFirstName("Nana")
	emp1Update.SetSalary(430_000)

	fmt.Printf("Emp[%d]Update Addres[%p] Value[%v]\n", 0, emp1, emp1)

	// 4. using slice employee value
	employeesValue := []employee.Employee{*emp1, *emp2, *emp3, *emp4, *emp5}

	for i, v := range employeesValue {
		fmt.Printf("Emp[%d] Addres[%p] Value[%v]\n", i, &v, v)
	}

	// 5. update emp1
	emp2Update := employeesValue[2]
	emp2Update.SetFirstName("Asep")
	emp2Update.SetSalary(150_000)

	fmt.Printf("Emp[%d]Update Addres[%p] Value[%v]\n", 2, emp3, emp3)

}