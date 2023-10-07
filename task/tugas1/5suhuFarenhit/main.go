package main

import (
	"fmt"
	"math"
)


//soal no 5
func main () {
	var celcius, fahrenheit float64
	fmt.Print("Mausukan angka untuk suhu celcius : ")
	fmt.Scan(&celcius)

	fahrenheit = (9 * celcius / 5) + 32
	fahrenheit = math.Round(fahrenheit * 100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

	fmt.Println("Hasil convert dari", celcius, " celcius menjadi satuan fahrenhait adalah :", fahrenheit, "derajat farhenhait")
}