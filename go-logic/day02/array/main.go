package main

import (
	"fmt"
)

func main() {
	// 1. declare arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	// 1.1 literal array
	numbers := [5]int{1, 2, 3, 4, 5}

	// 1.3 using three periods
	names := [...]string{"drama", "movies", "series", "short"}

	// 2, get value from array
	fmt.Println(numbers[1])
	fmt.Println("name : ", names[0])

	// 3. loping array
	for i := 0; i < len(names); i++ {
		fmt.Print(names[i], " ")
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}

	fmt.Println()

	for _, v := range numbers {
		fmt.Print(v)
	}

	for index, v := range numbers {
		fmt.Print(index, v)
	}

	fmt.Println()

	for _, v := range names {
		fmt.Print(v, " ")
	}

	for index, v := range names {
		fmt.Printf("%d.%s ", index, v)
	}

	println()

	// 4. comparing array
	movies := [3] string {"demonslayer", "spirit away", "madmax"}
	drama := [3] string {"got", "boys", "rome"}

	same := movies == drama

	fmt.Println(same)

	// 5. set key array for value
	score1 := [3] int {0:1, 1:2, 2:0}
	score2 := [3] int {0:2, 1:1, 2:0}

	sameScore1 := score1[2] == score2[2]
	sameScore2 := score1[0] == score2[1]
	sameScore3 := score1[1] == score2[2]
	fmt.Println(sameScore1)
	fmt.Println(sameScore2)
	fmt.Println(sameScore3)

	angka := [100] int {0:1, 99:5}
	fmt.Println(angka[0])
	fmt.Println(angka[99])
	fmt.Println(angka[67])

	// intial arrays with zero value
	var emptyFruits [5]string
	fmt.Println("Empty array : ", emptyFruits)

	// intial arrays with value 
	fruits := [5]string{"apel", "jeruk", "tomat", "durian", "mangga"}
	fmt.Println("Index\tValue\tLength")
	
	for i := 0; i < len(fruits); i++ {
		value := fruits[i]
		length := len(value)
		
		fmt.Printf("%d\t%s\t%d\n", i, value, length)
	}

	fmt.Printf("\nMemory adress arrays: %p\n", &fruits)
}