package main

import (
	"fmt"
)

func main() {
	var bilanganAngka int

	fmt.Printf("Bilangan angka : ")
	fmt.Scan(&bilanganAngka)

	if bilanganAngka%2 != 0 {
		fmt.Println("ini adalah bilangan ganjil TRUE")
	} else  {
		fmt.Println("ini adalah bilangan genap false")
	}
}
