package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari, luasPermukaan float64
	const phi float64 = 3.14

	fmt.Printf("Jari-jari bola : ")
	fmt.Scan(&jariJari)

	luasPermukaan = 4 * phi * jariJari * jariJari
	luasPermukaan = math.Round(luasPermukaan*10) / 10

	fmt.Println("Luas permukaan bola : ", luasPermukaan)
}
