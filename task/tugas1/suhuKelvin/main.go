package main 

import (
	"fmt"
	"math"
)

//soal no 6&7 sama
func main () {
	var celcius, kelvin float64
	fmt.Print("Masukan suhu dalam celcius : ")
	fmt.Scan(&celcius)

	kelvin = celcius + 273.15
	kelvin = math.Round(kelvin * 10) / 10

	fmt.Println("Hasil convert menjadi satuan Kelvin adalah :", kelvin, "derajat")
}