package main

import "fmt"

func initMatrix(matrix [5][5]int) [5][5]int {
	counter := 1
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = counter
			counter++
		}
	}
	return matrix
}

func displayMatrix(matrix [5][5]int)  {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%3d", matrix[row][col])
		}
		fmt.Println()
	}
}

func displayMatrixString(matrix [5][5]string)  {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%3s", matrix[row][col])
		}
		fmt.Println()
	}
}

func main() {
	// 1. declare matrix
	var matrix1 [3][3]int
	matrix1[0][0] = 10
	matrix1[0][1] = 15
	matrix1[0][2] = 20

	matrix1[1][0] = 6
	matrix1[1][1] = 7
	matrix1[1][2] = 8

	matrix1[2][0] = 11
	matrix1[2][1] = 12
	matrix1[2][2] = 13

	fmt.Println(matrix1[0])
	fmt.Println(matrix1[1])
	fmt.Println(matrix1[2])

	// 2. init matrix
	var matrix [5][5]int
	matrixCounter:= initMatrix(matrix)
	// fmt.Println(matrixCounter)
	displayMatrix(matrixCounter)

	var matrixStart [5][5]string
	matrixStart = matrixHollow(matrixStart)
	displayMatrixString(matrixStart)
}