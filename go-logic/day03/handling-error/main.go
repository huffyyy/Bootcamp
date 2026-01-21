package main

import (
	"errors"
	"fmt"
)

func addToCart(stock *int, total int) (int, error) {
	// check apakah jumlah total ditambahkan valid
	if total <= 0 {
		return 0, errors.New("Jumlah item harus lebih besar dari 0")
	}

	// check apakah stock nya avaible
	if *stock < total {
		return 0, errors.New("Stock tidak cukup untuk ditambahkan ke keranjang")
	}

	// mengurangi stock dan return total item yang ditambahkan
	*stock -= total
	return total, nil

}

func main() {

	// 1. initial stock & totalItem
	stok := 100
	totalItem := 0

	// add 3 item to cart 
	newTotal, err := addToCart(&stok, 3)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		totalItem += newTotal
		fmt.Printf("Succes add %d item, stock sisa : %d, total cart : %d", newTotal, stok, totalItem)
		fmt.Println()
	}

	// if total item value = negatif
	newTotal, err = addToCart(&stok, -1)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		totalItem += newTotal
		fmt.Printf("Succes add %d item, stock sisa : %d, total cart : %d", newTotal, stok, totalItem)
	}

	// if total item value = negatif
	newTotal, err = addToCart(&stok, 200)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		totalItem += newTotal
		fmt.Printf("Succes add %d item, stock sisa : %d, total cart : %d", newTotal, stok, totalItem)
	}
}