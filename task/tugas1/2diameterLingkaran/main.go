package main

import (
    "fmt"
    "math"
)


//soal no 2
func main() {
    const phi float64 = 3.14
    var luasLingkaran float64

    fmt.Print("diketahui luas lingkaran : ")
    fmt.Scan(&luasLingkaran)

    var diameter float64 = 2 * math.Sqrt(luasLingkaran/phi)
    diameter = math.Round(diameter*100) / 100 //handle if the input is bilanganya terlalu banyak koma, jadi di handle biar hanya ada 2 saja di output

    fmt.Println("Jadi, dengan luas lingkaran", luasLingkaran, ", diameter lingkarannya adalah", diameter)
}