package main

import (
	"fmt"
)


// soal no 15
func main() {
	var bilangan int
	var Hasil bool

	fmt.Print("Masukkan sebuah bilangan bulat : ")
	fmt.Scan(&bilangan)

	if bilangan == 0 {
		Hasil = true
		fmt.Println(Hasil)
	} else {
		Hasil = false
		fmt.Println(Hasil)
	}
}