package main 

import (
	"fmt"
	"math"
)

//soal no 6 & 7 sama
func main () {
	var celcius, kelvin float64
	fmt.Print("Masukan suhu dalam celcius : ")
	fmt.Scan(&celcius)

	kelvin = celcius + 273.15
	kelvin = math.Round(kelvin * 100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

	fmt.Println("Hasil convert dari", celcius, " celcius menjadi satuan kelvin adalah :", kelvin, "derajat kelvin")
}