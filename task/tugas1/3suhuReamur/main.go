package main

import (
	"fmt"
	"math"
)


//soal no 3
func main() {
	var celcius, reamur float64
	fmt.Print("Masukan angka untuk suhu celcius : ")
	fmt.Scan(&celcius)

	reamur = 4 * celcius / 5
	reamur = math.Round(reamur*100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

	fmt.Println("Hasil convert dari", celcius, " celcius menjadi satuan reamur adalah :", reamur, "derajat reamur")
}

