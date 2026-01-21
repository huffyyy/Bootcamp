package main

import "fmt"

func main() {
	// 1. declare variable immutable
	// data origin stock tidak bisa diubah
	stock := 100
	cart1 := 50
	fmt.Printf("stock: [value : %v\t address:%p]\n", stock, &stock )
	fmt.Printf("cart1: [value : %v\t address:%p]\n", cart1, &cart1 )

	cart2 := stock - 25
	fmt.Printf("cart2: [value : %v\t address:%p]\n", cart2, &cart2 )
	fmt.Printf("stock: [value : %v\t address:%p]\n", stock, &stock )

	// 2. pointer
	stockLaptop := 150
	cartMe := &stockLaptop
	fmt.Printf("stockLaptop: [value : %v\t address:%p]\n", stockLaptop, &stockLaptop )
	fmt.Printf("cartMe: [value : %v\t address:%p]\n", cartMe, &cartMe )

	fmt.Printf("cartMe: [value : %v\t address:%p]\n", *cartMe, &cartMe )

	// 3. modify pointer
	cart3 := &stock
	
	*cart3 = *cart3 - 10
	fmt.Printf("stock: [value : %v\t address:%p]\n", stock, &stock )
	fmt.Printf("cart3: [value : %v\t address:%p]\n", cart3, &cart3 )
	
	// cart4 := &cart3
	// *cart4 = *cart3 - 20

	sisaStock := addStock(&stock, 2)
	fmt.Printf("sisa stock: [value : %v\t address:%p]\n", sisaStock, &sisaStock )
}

func addStock(stock *int, buy int) int  {
	*stock = *stock - buy
	return  *stock
}