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

	luasLingkaran := phi * jariJari * jariJari
	luasLingkaran = math.Round(luasLingkaran*10)/10

	fmt.Println("Jadi, Luas lingkaran adalah : ", luasLingkaran)
}



