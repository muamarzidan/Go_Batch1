package main

import (
	"fmt"
	"math"
)


//soal no 1
func main() {
	var jariJari float64
	const phi float64 = 3.14

	fmt.Print("Jari-jari lingkaran : ")
	fmt.Scan(&jariJari)

	r := jariJari

	luasLingkaran := phi * r * r
	luasLingkaran = math.Round(luasLingkaran*100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

	fmt.Println("Jadi, dengan jari-jari", jariJari, ", Luas lingkaran adalah:", luasLingkaran)
}



