package main

import (
	"fmt"
	"strings"
)


//soal no 13 & 14 sama
func main() {
	var hurufKonsonan string
	var Hasil bool

	fmt.Print("Masukkan sebuah huruf : ")
	fmt.Scan(&hurufKonsonan)
	hurufKonsonan = strings.ToLower(hurufKonsonan)

	if hurufKonsonan != "a" && hurufKonsonan != "i" && hurufKonsonan != "u" && hurufKonsonan != "e" && hurufKonsonan != "o" {
		Hasil = true
		fmt.Println(Hasil)
	} else {
		Hasil = false
		fmt.Println(Hasil)
	}
}