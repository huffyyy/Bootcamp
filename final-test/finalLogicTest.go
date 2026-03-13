package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1
	numbers := []int{2, 3, 5, 7, 11, 13}

	result1 := twoSum(numbers, 9)
	fmt.Println(result1)

	result2 := twoSum(numbers, 12)
	fmt.Println(result2)

	// 2
	fmt.Println(palindrome("MalaM"))
	fmt.Println(palindrome("levEl"))
	fmt.Println(palindrome("kaSur ini ruSak"))
	fmt.Println(palindrome("saYur"))

	// 3
	fmt.Println(countString("xhixhix", "x"))
	fmt.Println(countString("xhixhix", "hi"))
	fmt.Println(countString("mic", "mic"))
	fmt.Println(countString("haha", "ho"))
	fmt.Println(countString("xxxxyz", "xx"))

	// 4
	fmt.Println(plusOne([]int{1, 3, 2, 4}))
	fmt.Println(plusOne([]int{1, 4, 8, 9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))

	// 5
	triangleNumber(7, 1)
	triangleNumber(7, 5)

}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{numbers[i], numbers[j]}
			}
		}
	}
	return nil
}

func palindrome(s string) bool {
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func countString(s string, target string) int {
	count := 0
	targetLen := len(target)

	for i := 0; i <= len(s)-targetLen; i++ {
		if s[i:i+targetLen] == target {
			count++
		}
	}
	return count
}

func plusOne(numbers []int) []int {
	for i := len(numbers) - 1; i >= 0; i-- {
		numbers[i]++
		if numbers[i] < 10 {
			return numbers
		}
		numbers[i] = 0
	}

	newResult := []int{1}
	for i := 0; i < len(numbers); i++ {
		newResult = append(newResult, 0)
	}
	return newResult
}

func triangleNumber(numbers int, start int) {
	for i := 0; i < numbers; i++ {
		n := start + i
		for j := 0; j <= i; j++ {
			fmt.Print(n, " ")
			n++
		}
		fmt.Println()
	}
}
