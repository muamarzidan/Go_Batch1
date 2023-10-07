package main

import (
	"fmt"
)

//soal no 9
func main() {
	var bilangan int

	fmt.Printf("Bilangan angka : ")
	fmt.Scan(&bilangan)

	if bilangan%2 != 0 {
		fmt.Println("ini adalah bilangan ganjil, TRUE")
	} else  {
		fmt.Println("ini adalah bilangan genap, FALSE")
	}
}
