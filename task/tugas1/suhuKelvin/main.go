package main 

import (
	"fmt"
	"math"
)

func main () {
	var celcius, kelvin float64
	fmt.Print("Suhu dalam celcius : ")
	fmt.Scan(&celcius)

	kelvin = celcius + 273.15
	kelvin = math.Round(kelvin * 10) / 10

	fmt.Println("Hasil convert : ", kelvin)
}