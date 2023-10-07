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
	kelilingPersegiPanjang = math.Round(kelilingPersegiPanjang * 1) / 1 //handle if the input is not a bilangan bulat, so the result is can be bilangan bulat

	fmt.Println("Jadi, dari panjang", panjang, "dan lebar", lebar, ", maka keliling persegi panjang adalah :", kelilingPersegiPanjang)
}
