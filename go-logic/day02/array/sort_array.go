package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. sorting int array
	numbers := [...]int{9, 6, 87, 0, 78, 54, 12, 7, 8}
	for _, v := range numbers {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// after sort 
	sort.Ints(numbers[:])
	for _, v := range numbers {
		fmt.Printf("%d ", v)
	}

	fmt.Println()

	// 2. sort string array
	movies := [...]string{"avenger", "saga", "furiosa", "demon", "boy", "madmax", "got"}
	sort.Strings(movies[:])
	for _, v := range movies {
		fmt.Printf("%s ",  v)
		fmt.Println()
	}
}