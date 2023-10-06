package main

import (
	"fmt"
	"math"
)

func main () {
	var celcius, farenhit float64
	fmt.Print("Suhu dalam celcius : ")
	fmt.Scan(&celcius)

	farenhit = (9 * celcius / 5) + 32
	farenhit = math.Round(farenhit * 10) / 10

	fmt.Println("Hasil convert : ", farenhit)
}