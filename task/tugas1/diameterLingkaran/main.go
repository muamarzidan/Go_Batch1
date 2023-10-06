package main 

import "fmt"

func diameterLingkaran()  {
	const phi float64 = 3.14
	var luasLingkaran float64

	fmt.Print("Masukkan luas lingkaran : ")
	fmt.Scan(&luasLingkaran)

	diameter := 2 * (luasLingkaran / phi)

	fmt.Println("Diameter lingkaran adalah", diameter)
}