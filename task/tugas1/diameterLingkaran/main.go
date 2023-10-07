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

    if luasLingkaran <= 0 {
        fmt.Println("Masukan luas lingkaran dengan angka yang benar")
        return
    }

    var diameter float64 = 2 * math.Sqrt(luasLingkaran/phi)
    diameter = math.Round(diameter*10)/10

    fmt.Println("Jadi, diameter lingkarannya adalah", diameter)
}