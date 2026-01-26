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

type manager struct {
	employee
	email string
	totalStaff int64
}

type programmer struct {
	employee
	placement string
}

func toString(e employee) string  {
	return fmt.Sprintf("FullName : %s %s, HireDate : %s, Salary : %2.f",
		e.firstName, e.lastName, e.hireDate.Format("2026-01-26"), e.salary)
}

func toStringMgr(m manager) string  {
	return fmt.Sprintf("FullName : %s %s, HireDate : %s, Salary : %2.f, Email : %s, TotalStaff : %d",
		m.firstName, m.lastName, m.hireDate.Format("2026-01-26"), m.salary, m.email, m.totalStaff)
}

func toStringProg(p programmer) string  {
	return fmt.Sprintf("FullName : %s %s, HireDate : %s, Salary : %2.f, Placement : %s",
		p.firstName, p.lastName, p.hireDate.Format("2026-01-26"), p.salary, p.placement)
}

func main() {

	// 1. create object manager
	mgr1 := manager{
		employee: employee{
			firstName: "Husnul",
			lastName: "Fikri",
			hireDate: time.Now(),
			salary: 200_000,
		},
		email: "husnul@mail.com",
		totalStaff: 10,
	}


	// 2. object programmer
	prog1 := programmer{
		employee: employee{
			firstName: "Alicia",
			lastName: "Juned",
			hireDate: time.Now(),
			salary: 250_000,
		},
		placement: "Internal Project",
	}
	fmt.Println(toString(mgr1.employee))
	fmt.Println(toString(prog1.employee))
	fmt.Println(toStringMgr(mgr1))
	fmt.Println(toStringProg(prog1))
}