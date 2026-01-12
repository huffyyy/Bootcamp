package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. declare with datatype
	var fullName string = "Bootcamp CodeId"
	fmt.Println(fullName)

	// 2. declare multiple variable datatype using var
	var (
		fistName string = "Jhone"
		lastName string = "Snow"
		salary float64 = 5_000_000 //digit separator
		hireDate time.Time = time.Now()
	)

	fmt.Println(fistName, lastName, salary, hireDate)

	// 3. declare variable tanpa datatype 
	var (
		departemetId = 10
		departemetName = "Finance"
		rateSalary = 45_00_00.654
	)

	fmt.Println(departemetId, departemetName, rateSalary)

	// 4. declare multiple variable in oneline with zero value
	var salary1, salary2, salary3 float64 //default value 0
	var myName string // default value "0"

	salary1 = 10.90
	myName = "Billy Kid"
	fmt.Println(myName, salary1, salary2, salary3)

	// 5. short variable
	totalSalary := 10_000_00.00
	taxPph := 0.2
	fmt.Println(totalSalary, taxPph)

	// 6. multiple varible with shortvariable
	isActive, customerName, orderDate, price := false, "PT.XYZ", time.Now(), 5_000_000
	fmt.Println(isActive, customerName, orderDate, price)

	// 7. constanta
	const pajak = 12.0
}