package main

import (
	"fmt"
	"strings"
)


//soal no 12
func main() {
	var hurufVokal string
	var Hasil bool

	fmt.Print("Masukkan sebuah huruf : ")
	fmt.Scan(&hurufVokal)
	hurufVokal = strings.ToLower(hurufVokal)

	if hurufVokal == "a" || hurufVokal == "i" || hurufVokal == "u" || hurufVokal == "e" || hurufVokal == "o" {
		Hasil = true
		fmt.Println(Hasil)
	} else {
		Hasil = false
		fmt.Println(Hasil)
	}
}