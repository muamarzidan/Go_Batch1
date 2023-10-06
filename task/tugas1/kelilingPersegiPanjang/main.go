package main

import "fmt"

func main() {
	var panjang, lebar float64

	fmt.Print("Panjang persegi panjang : ")
	fmt.Scan(&panjang)

	fmt.Print("Lebar persegi panjang : ")
	fmt.Scan(&lebar)

	var kelilingPersegiPanjang float64  = 2 * (panjang + lebar)

	fmt.Println("Keliling persegi panjang : ", kelilingPersegiPanjang)
}
