package main

import "fmt"

func boxHollow (n int) {
	for i := 0; i < n; i++ { // outer loop u/pindah baris
		for j := 0; j < n; j++ { // inner loop u/print
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Printf("%s", "*")
			} else {
				fmt.Printf("%s", " ")
			}
		}
		fmt.Println()
	}
}

func triangle01(n int) {
	for i := 0; i < n; i++ {      
		for j := 0; j <= i; j++ { 
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func triangle02(n int) {
    for i := 0; i < n; i++ {
        for j := 0; j < n-i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

func triangle05(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				fmt.Print(i)
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func triangle06(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				fmt.Print(i + 1)
			} else if i == n-1 {
				fmt.Print(n + j)
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n % i == 0 {
			return false
		}
	}	
	return true
}

func boxStar (n int) {
	for i := 0; i < n; i++ { // outer loop u/pindah baris
		for j := 0; j < n; j++ { // inner loop u/print
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func countDigits(n int) int {
	counter := 0
	sisa := n
	for {
		sisa = sisa/10
		counter++
		if sisa == 0 {
			break
		}
	}
	return counter
}

func main() {
	n := 5
	boxStar(n)
	boxHollow(n)
	triangle01(n)
	triangle02(n)
	triangle05(n)
	triangle06(n)

	countDigit := countDigits(3452233)
	fmt.Println("countdigit : ", countDigit)

	fmt.Println("countdigit : ", countDigits(34255))

	fmt.Println("is prime : ", isPrime(5))
	fmt.Println("is prime : ", isPrime(6))

}
