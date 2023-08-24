package main

import (
	"fmt"
)

func Kelipatan7(number int) bool {
	return number%7 == 0
}

func main() {
	var number int
	fmt.Print("======= Program Menentukan Bilangan Kelipatan 7 =======\n")
	fmt.Print("-------------------------------------------------------\n")
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&number)
	fmt.Print("-------------------------------------------------------\n")

	if Kelipatan7(number) {
		fmt.Printf("%d adalah kelipatan dari 7.\n", number)
	} else {
		fmt.Printf("%d bukan kelipatan dari 7.\n", number)
	}
}
