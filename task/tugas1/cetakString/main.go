package main

import (
	"fmt"
)


func main() {
	var nama string

	fmt.Print("Masukkan kata : ")
	fmt.Scan(&nama)

	fmt.Println("Hai " + nama + ", saya sebenarnya adalah manusia :)")
}