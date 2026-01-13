package main

import (
	"fmt"
	"math/rand"
)

func random(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findNumber(arr [5]int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func fillArray(n int) [5]int {
	var arr [5]int
	for i := 0; i < n && i < 5; i++ {
		arr[i] = i
	}
	return arr
}

func fillArrayRandom(n int) [5]int {
	var arr [5]int
	for i := 0; i < n && i < 5; i++ {
		arr[i] = random(1, 20)
	}
	return arr
}

func fillArrayPrime(n int) [5]int {
	var arr [5]int
	count := 0
	num := 2
	for count < 5 && count < n {
		if isPrime(num) {
			arr[count] = num
			count++
		}
		num++
	}
	return arr
}

func displayArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func main() {
	// 1. store array in variable then display
	array1 := fillArray(5)
	displayArray(array1)

	array2 := fillArrayRandom(5)
	displayArray(array2)

	// 2. passing func fillArray to display array without variables
	displayArray(fillArray(5))
	displayArray(fillArrayRandom(5))
	displayArray(fillArrayPrime(5))

	// 3. find element array
	fmt.Printf("found number index : %d\n", findNumber(fillArrayRandom(5), 10))
}