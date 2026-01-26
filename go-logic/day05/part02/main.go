package main

import (
	"fmt"
	"time"
)

type location struct {
	id int64
	address string
}

type departement struct {
	id int64
	departementName string
	location location
}

type employee struct {
	firstName string
	lastName  string
	hireDate  time.Time
	salary float64
	departement departement
}

func toString(e employee) string  {
	return fmt.Sprintf("FullName : %s %s, HireDate : %s, Salary : %2.f, Departement : %s, Address : %s ", 
		e.firstName, e.lastName, e.hireDate.Format("2026-01-26"), e.salary, e.departement.departementName, e.departement.location.address)
}

func main() {

	// 1. literal
	emp1 := employee {
		firstName: "Husnul",
		lastName: "Fikri",
		hireDate: time.Now(),
		salary: 150_000,
		departement: departement{
			id: 10,
			departementName: "IT",
			location: location{
				id : 1001,
				address: "Jakarta",
			},
		},
	}

	fmt.Println(toString(emp1))
}