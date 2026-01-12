package main

import "fmt"

// Quiz Time (1)
func findDivisor(n int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func extractDigit(n int) {
	for n > 0 {
		fmt.Print(n%10, " ")
		n = n / 10
	}
	fmt.Println()
}

// Quiz Time (2)
func triangle01(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func triangle02(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= n-i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func numberPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := n - i; j >= 1; j-- {
			fmt.Print(j, " ")
		}
		for j := 2; j <= n-i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

// Quiz Time (3)
func numberPattern01(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%2 == 1 {
				fmt.Print(i, " ")
			} else {
				fmt.Print(n-i+1, " ")
			}
		}
		fmt.Println()
	}
}

func numberPattern02(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if (i+j)%2 != 0 {
                fmt.Print(j, " ")
            } else {
                fmt.Print("- ")
            }
        }
        fmt.Println()
    }
}

// Quiz Time (4)
func isPalindrome(text string) bool {
	clean := ""
	for _, ch := range text {
		if ch != ' ' {
			if ch >= 'A' && ch <= 'Z' {
				ch = ch + 32 
			}
			clean += string(ch)
		}
	}
	length := len(clean)
	for i := 0; i < length/2; i++ {
		if clean[i] != clean[length-1-i] {
			return false
		}
	}
	return true
}

// Quiz Time (5)
func reverse(text string) string {
	result := ""
	for i := len(text) - 1; i >= 0; i-- {
		result += string(text[i])
	}
	return result
}


func checkBraces(text string) bool {
	count := 0
	for i := 0; i < len(text); i++ {
		if text[i] == '(' {
			count++
		} else if text[i] == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

// Quiz Time (6)
func isNumberPalindrome(n int) bool {
	original := n
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + (n % 10)
		n = n / 10
	}
	return original == reversed
}


func main() {
	findDivisor(6)
	findDivisor(24)
	findDivisor(7)

	extractDigit(12234)
	extractDigit(5432)
	extractDigit(1278)

	triangle01(5)
	triangle02(5)
	numberPyramid(8)

	numberPattern01(9)
	numberPattern01(5)

	numberPattern02(9)
	numberPattern02(5)

	fmt.Println(isPalindrome("Kasur ini rusak"))
	fmt.Println(isPalindrome("tamaT")) 
	fmt.Println(isPalindrome("Aku Usa"))

	fmt.Println(reverse("ABCD"))
	fmt.Println(reverse("tamaT"))
	fmt.Println(reverse("XYnb"))

	fmt.Println(checkBraces("(())"))
	fmt.Println(checkBraces("()()"))
	fmt.Println(checkBraces("((()"))
	fmt.Println(checkBraces("(()))((())")) 
	
	fmt.Println(isNumberPalindrome(121)) 
	fmt.Println(isNumberPalindrome(2147447412))
	fmt.Println(isNumberPalindrome(333))        
	fmt.Println(isNumberPalindrome(110))
	fmt.Println(isNumberPalindrome(11))
}