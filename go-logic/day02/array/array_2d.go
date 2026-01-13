package main

import (
	"fmt"
	"math/rand"
)

func isPrime(n int) bool {
	if n < 2 { return false }
	for i := 2; i*i <= n; i++ {
		if n%i == 0 { return false }
	}
	return true
}

func fillArrayWithRandom(matrix [5][5]int) [5][5]int {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			matrix[row][col] = rand.Intn(11) 
		}
	}
	return matrix
}

func fillArrayWithPrime(matrix [5][5]int) [5][5]int {
	num := 2
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			for !isPrime(num) {
				num++
			}
			matrix[row][col] = num
			num++
		}
	}
	return matrix
}

func displayMatrix(matrix [5][5]int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println() 
	}
	fmt.Println()
}

func main() {
	// 1. fill matrix array manually
	var matrix1 [3][3]int
	matrix1[0][0] = 10
	matrix1[0][1] = 15
	matrix1[0][2] = 20
	fmt.Println("Baris 0 Matrix 1:", matrix1[0])

	// 2. fill matrix array literally
	matrix2 := [3][2]int{
		{1, 2}, // Baris 0
		{3, 4}, // Baris 1
		{5, 6}, // Baris 2
	}
	fmt.Println("Baris 0 Matrix 2:", matrix2[0])

	// 3. fill matrix array with random number
	var matrix [5][5]int
	
	fmt.Println("\nMatrix Random:")
	matrix = fillArrayWithRandom(matrix)
	displayMatrix(matrix)

	// 4. matrix array with prime number 
	fmt.Println("Matrix Prime:")
	matrixPrime := fillArrayWithPrime(matrix)
	displayMatrix(matrixPrime)
}