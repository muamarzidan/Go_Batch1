package main

import "fmt"

func main() {
	var total, jumlahKali, input int

	for total <= 100 {
		fmt.Print("Masukkan angka: ")
		fmt.Scanln(&input)
		total += input
		jumlahKali++
	}
}
