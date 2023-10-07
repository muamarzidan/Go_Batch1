package main

import (
	"fmt"
)

//soal no 9
func main() {
	var bilangan int
	var Hasil bool

	fmt.Print("Bilangan bilangan bulat : ")
	fmt.Scan(&bilangan)

	if bilangan%2 != 0 {
		Hasil = true
		fmt.Println(Hasil)
	} else  {
		Hasil = false
		fmt.Println(Hasil)
	}
}
