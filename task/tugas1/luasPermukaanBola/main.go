package main

import (
	"fmt"
	"math"
)


//soal no 8
func main() {
	var jariJari, luasPermukaan float64
    const phi float64 = 3.14

    fmt.Printf("Jari-jari bola : ")
    fmt.Scan(&jariJari)

    r := jariJari

    luasPermukaan = 4 * phi * r * r
    luasPermukaan = math.Round(luasPermukaan*10) / 10

    fmt.Println("Luas permukaan bola : ", luasPermukaan)
}
