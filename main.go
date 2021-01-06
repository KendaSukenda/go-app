package main

import (
	"fmt"
	goModule "github.com/KendaSukenda/go-module"
)

func main() {
	perkalian := goModule.Perkalian(2, 3)
	pengurangan := goModule.Pengurangan(2, 3)
	penjumlahan := goModule.Penjumlahan(2, 3)
	pembagian := goModule.Pembagian(2, 3)
	modulo := goModule.Modulo(2, 3)

	fmt.Println("Hasil perkalian ", perkalian)
	fmt.Println(pengurangan)
	fmt.Println(penjumlahan)
	fmt.Println(pembagian)
	fmt.Println(modulo)
}
