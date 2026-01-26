package main

import (
	"fmt"
	"strings"
)

// Quiz Time (1)
func upperCaseExcept(words []string, except string) []string {
	for i := 0; i < len(words); i++ {
		if words[i] != except {
			words[i] = strings.ToUpper(words[i])
		}
	}
	return words
}

func findMinMax(numbers []uint) []uint {
	if len(numbers) == 0 {
		return []uint{}
	}
	min, max := numbers[0], numbers[0]

	for _, v := range numbers {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return []uint{min, max}
}

// Quiz Time (2)
func findMinRange(numbers []int, start int, end int) []int {
	if start < 0 { start = 0 }
	if end > len(numbers) { end = len(numbers) }
	minVal := numbers[start]
	indexPos := start
	for i := start; i < end; i++ {
		if numbers[i] < minVal {
			minVal = numbers[i]
			indexPos = i
		}
	}
	return []int{minVal, indexPos}
}
func findMaxRange(numbers []int, start int, end int) []int {
	if start < 0 { start = 0 }
	if end > len(numbers) { end = len(numbers) }
	maxVal := numbers[start]
	indexPos := start
	for i := start; i < end; i++ {
		if numbers[i] > maxVal {
			maxVal = numbers[i]
			indexPos = i
		}
	}
	return []int{maxVal, indexPos}
}

func evenOddOrder(numbers []int) []int {
	var evens []int
	var odds []int
	for _, v := range numbers {
		if v%2 == 0 {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	return append(evens, odds...)
}

// Quiz Time (3)
func rotateArray(nums []int, n int) []int {
	if len(nums) == 0 {
		return nums
	}
	n = n % len(nums)

	for i := 0; i < n; i++ {
		firstElement := nums[0]
		copy(nums, nums[1:])
		nums[len(nums)-1] = firstElement
	}

	return nums
}

// Quiz Time (4)
func matrixDiagonal01(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				fmt.Printf("%2d ", i+1) 
			} else if j > i {
				fmt.Printf("%2d ", 10)
			} else {
				fmt.Printf("%2d ", 20)
			}
		}
		fmt.Println() 
	}
}

func matrixDiagonal02(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				fmt.Printf("%2d ", i+1) 
			} else if j < i {
				fmt.Printf("%2d ", 10)
			} else {
				fmt.Printf("%2d ", 20)
			}
		}
		fmt.Println() 
	}
}

// Quiz Time (5)
func boxHollowNumbers(n int) {
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i == 0 {
                fmt.Printf("%2d ", j)
            } else if i == n-1 {
                fmt.Printf("%2d ", (n-1)+j)
            } else if j == 0 {
                fmt.Printf("%2d ", i)
            } else if j == n-1 {
                fmt.Printf("%2d ", (n-1)+i)
            } else {
                fmt.Printf("   ") 
            }
        }
        fmt.Println()
    }
}

func matrixBoxNumbers(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i == n && j == n {
				fmt.Printf("%2d ", (n-1)*n)
			} else if i == n {
				fmt.Printf("%2d ", (j+3)*n)
			} else if j == n {
				fmt.Printf("%2d ", (i+3)*n)
			} else {
				fmt.Printf("%2d ", i+j)
			}
		}
		fmt.Println()
	}
}


func main() {
	fmt.Println(upperCaseExcept([]string{"code", "java", "cool"}, "java"))
	fmt.Println(upperCaseExcept([]string{"black", "pink", "venom"}, "venom"))

	res := findMinMax([]uint{2, 3, 4, 5, 6, 7, 8, 9, 1, 10})
	fmt.Printf("[%d, %d]\n", res[0], res[1])

	resMin := findMinRange([]int{5, 3, 4, 2, 6, 7, 8, 9, 1, 10}, 0, 10)
	fmt.Printf("[%d, %d]\n", resMin[0], resMin[1]) 
	resMin2 := findMinRange([]int{5, 3, 4, 2, 6, 7, 8, 9, 1, 10}, 0, 7)
	fmt.Printf("[%d, %d]\n", resMin2[0], resMin2[1]) 

	resMax := findMaxRange([]int{1, 22, 3, 4, 5, 10, 7, 8, 9, 49}, 0, 10)
	fmt.Printf("[%d, %d]\n", resMax[0], resMax[1])
	resMax2 := findMaxRange([]int{1, 22, 3, 4, 5, 10, 7, 8, 9, 49}, 2, 7)
	fmt.Printf("[%d, %d]\n", resMax2[0], resMax2[1])

	fmt.Println(evenOddOrder([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	array := []int{12, 15, 1, 5, 20}

	fmt.Println("Original :", array)
	fmt.Println("1 Rotate :", rotateArray([]int{12, 15, 1, 5, 20}, 1))
	fmt.Println("2 Rotate :", rotateArray([]int{12, 15, 1, 5, 20}, 2))
	fmt.Println("3 Rotate :", rotateArray([]int{12, 15, 1, 5, 20}, 3))

	matrixDiagonal01(5)
	matrixDiagonal02(5)

	boxHollowNumbers(7)
	matrixBoxNumbers(7)
}