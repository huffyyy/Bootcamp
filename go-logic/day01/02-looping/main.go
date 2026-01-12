package main

import (
	"fmt"
)

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

func triangleX(n int) {
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            if i == j || i + j == n + 1 {
                fmt.Print(i)
            } else {
                fmt.Print(" ") 
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

func okYes(n int) {
	for i := 1; i < n; i++ {
		if i % 3 == 0 && i % 4 == 0 {
			fmt.Println("OkYes")
		} else if i % 3 == 0 {
			fmt.Println("Ok")
		} else if i % 4 == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println(i)
		}
	}
}

func sumOk(n int) int  {
	sum := 0

	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 4 != 0 {
			sum+=i
		}
	}
	return sum
}

func counOk(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 4 != 0 {
			count++
		}
	}
	return count
}

func sumOkContinue(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 4 == 0  {
			break
		} else if i % 3 == 0 {
			sum += i
		} else if i % 4 == 0 {
			continue
		} else {
			continue
		}
	}
	return sum
}

	func countVowel1(word string) (int, int) {
		vocal, consonant := 0, 0
		for i := 0; i < len(word); i++ {
			if word[i] == 'a' || word[i] == 'i'|| word[i] == 'u' || word[i] =='e' || word[i] == 'o' {
				vocal++
			} else if word[i] != ' ' {
				consonant++
			}
		}
		return vocal, consonant
	}

	func countVowel2(word string) (int, int) {
		vocal, consonant := 0, 0		
		for _, v := range word {
			if v == 'a' || v == 'i' || v == 'u' || v == 'e' || v == 'o'  {
				vocal++
			} else if v != ' ' {
				consonant++
			}
		}
		return vocal, consonant
	}


func main() {
	n := 5
	boxStar(n)
	boxHollow(n)
	triangle01(n)
	triangle02(n)
	triangle05(n)
	triangle06(n)
	triangleX(n)

	countDigit := countDigits(3452233)
	fmt.Println("countdigit : ", countDigit)

	fmt.Println("countdigit : ", countDigits(34255))

	fmt.Println("is prime : ", isPrime(5))
	fmt.Println("is prime : ", isPrime(6))

	okYes(15)
	fmt.Println("SumOk : ", sumOk(15))
	fmt.Println("CountOk : ", counOk(15))

	fmt.Println("SumOkContinue : ", sumOkContinue(25))

	// goto statement // label statement
	counter := 0
	target: 
		counter++
		fmt.Println("Counter : ", counter)
		if counter < 5 {
			goto target 
		}
		fmt.Println("Selesai!")
	
	intro := "Hello coding bootcamp codeid"
	fmt.Println(intro)
	
	vocal, consonant := countVowel1(intro)
	fmt.Println("Vocal : ", vocal)
	fmt.Println("Consonant : ", consonant)

	vocal, consonant = countVowel2(intro)
	fmt.Println("Vocal : ", vocal)
	fmt.Println("Consonant : ", consonant)

	// multilateral string
	str1 := `
			bootcamp codeid
			golang batch#30
	`
	fmt.Println(str1)

	str2 := "Folder in \"c:\\bootcamp\\go\\\""
	fmt.Println(str2)

	// string to byte
	username := "tanjiro"
	for i := 0; i < len(username); i++ {
		fmt.Println(username[i] ," ")
	}

	// byte to string
	for i := 0; i < len(username); i++ {
		fmt.Println(string(username[i]), " ")
	}

	// using range 
	for i, v := range username {
		fmt.Printf("%d:%v ", i, string(v))
	}

	// username to rune
	runeUsername := []rune(username)
	for i := 0; i < len(username); i++ {
		fmt.Printf("%d", runeUsername[i])
	}

	// rune literals
var (
		tab       = '\t'
		newline   = '\n'
		backspace = '\b'
		hagul     = '가'  
		arab      = 'ع'    
		hexa      = '\xFF' 
		unicode   = '\u0369' 
	)

	fmt.Println("=== Output Angka (Rune ID) ===")
	fmt.Println("tab : ", tab)
	fmt.Println("newline : ", newline)
	fmt.Println("backspace : ", backspace)
	fmt.Println("hagul : ", hagul)
	fmt.Println("arab : ", arab)
	fmt.Println("hexa : ", hexa)
	fmt.Println("unicode : ", unicode)
}
