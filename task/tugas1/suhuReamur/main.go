package main

import "fmt"

func main() {
	var celcius, reamur float64
	fmt.Print("Masukan suhu dalam celcius : ")
	fmt.Scan(&celcius)

	reamur = 4 * celcius / 5
	fmt.Println("Hasil convert : ", reamur)
}
