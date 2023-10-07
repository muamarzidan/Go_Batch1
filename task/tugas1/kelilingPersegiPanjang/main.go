package main

import (
	"fmt"
	"math"
)


//soal no 4
func main() {
	var panjang, lebar float64

	fmt.Print("Panjang persegi panjang : ")
	fmt.Scan(&panjang)

	fmt.Print("Lebar persegi panjang : ")
	fmt.Scan(&lebar)

	var kelilingPersegiPanjang float64  = 2 * (panjang + lebar)
	kelilingPersegiPanjang = math.Round(kelilingPersegiPanjang * 10) / 10

	fmt.Println("Jadi, keliling persegi panjang : ", kelilingPersegiPanjang)
}
