package main

import "fmt"


//soal no 3
func main() {
	var celcius, reamur float64
	fmt.Print("Masukan suhu dalam celcius : ")
	fmt.Scan(&celcius)

	reamur = 4 * celcius / 5
	fmt.Println("Hasil convert menjadi satuan reamur adalah :", reamur, "derajat")
}
