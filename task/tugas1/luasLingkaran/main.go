package main

import "fmt"

func main() {
	var jariJari float64
	const phi float64 = 3.14

	fmt.Print("Masukkan jari-jari lingkaran : ")
	fmt.Scan(&jariJari)

	luasLingkaran := phi * jariJari * jariJari

	fmt.Println("Luas lingkaran adalah", luasLingkaran)
}



