package main

import "fmt"

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
	matrixDiagonal01(5)
	matrixDiagonal02(5)

	boxHollowNumbers(7)
	matrixBoxNumbers(7)
}