package main

import (
	"fmt"
)

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	// rotateArray(numbers)
	result := (rotateArray(numbers))
	fmt.Println(result)
}

func rotateArray(numbers []int) []int {
	firstElement := numbers[0]
	for i := 0; i < len(numbers)-1; i++ {
		numbers[i] = numbers[i+1]
	}
	numbers[len(numbers)-1] = firstElement
	return numbers
}
