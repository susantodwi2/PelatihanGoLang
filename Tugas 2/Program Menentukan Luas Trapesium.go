package main

import (
	"fmt"
)

func LuasTrapesium(sisi1, sisi2, tinggi float64) float64 {
	return 0.5 * (sisi1 + sisi2) * tinggi
}

func main() {
	var sisi1, sisi2, tinggi float64
	fmt.Print("======= Program Menentukan Luas Trapesium =======\n")
	fmt.Print("-------------------------------------------------\n")
	fmt.Print("Masukkan panjang Sisi atas: ")
	fmt.Scan(&sisi1)
	fmt.Print("Masukkan panjang Sisi bawah: ")
	fmt.Scan(&sisi2)
	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scan(&tinggi)
	fmt.Print("\n-------------------------------------------------\n")

	area := LuasTrapesium(sisi1, sisi2, tinggi)
	fmt.Printf("Luas trapesium adalah: %.2f\n", area)
}
