package main

import (
	"fmt"
)


//soal no 16
func main() {
	var namaSaya string

	fmt.Print("Masukkan kata : ")
	fmt.Scan(&namaSaya)

	fmt.Println("Hai, " + namaSaya + ", saya sebenarnya adalah manusia :)")
}