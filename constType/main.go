package main

import "fmt"

// func main() {
// 	const pi float64 = 3.14
// 	var jariJari float64 = 5

// 	var luasLingkaran float64 = pi * (jariJari * jariJari)
// 	var kelilingLingkaran float64 = 2 * pi * jariJari

// 	fmt.Println("Luas Lingkaran: ", luasLingkaran)
// 	fmt.Println("Keliling Lingkaran: ", kelilingLingkaran)
// }

// func constInt() {
// 	const MAXTINGGI int = 170
// 	var tinggi int
// 	tinggi = 165
// 	fmt.Println(MAXTINGGI - tinggi)
// }

func main() {
	const pi float64 = 3.14
	var jariJari float64
	fmt.Scanf("%f", &jariJari)

	var volume float64 = 4 / 3 * pi * (jariJari * jariJari * jariJari)

	fmt.Println("Volume Bola: ", volume)
}

