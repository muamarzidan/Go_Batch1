package main

import (
	"fmt"
	"math"
)


//soal no 8
func main() {
	var jariJari, luasPermukaan float64
    const phi float64 = 3.14

    fmt.Print("Jari-jari bola : ")
    fmt.Scan(&jariJari)

    r := jariJari

    luasPermukaan = 4 * phi * r * r
    luasPermukaan = math.Round(luasPermukaan*100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

    fmt.Println("Jadi, dari jari-jari", r, "Luas permukaan bola : ", luasPermukaan)
}
