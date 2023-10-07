package main

import (
	"fmt"
	"math"
)


//soal no 5
func main () {
	var celcius, fahrenheit float64
	fmt.Print("Mausukan suhu dalam celcius : ")
	fmt.Scan(&celcius)

	fahrenheit = (9 * celcius / 5) + 32
	fahrenheit = math.Round(fahrenheit * 10) / 10

	fmt.Println("Hasil convert menjadi satuan fahrenheit adalah :", fahrenheit, "derajat")
} 