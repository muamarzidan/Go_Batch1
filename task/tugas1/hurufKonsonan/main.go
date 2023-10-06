package main

import (
	"fmt"
	"strings"
)


//soal no 13&14 sama
func main() {
	var hurufKonsonan string

	fmt.Print("Masukkan huruf : ")
	fmt.Scan(&hurufKonsonan)
	hurufKonsonan = strings.ToLower(hurufKonsonan)

	if hurufKonsonan != "a" && hurufKonsonan != "e" && hurufKonsonan != "i" && hurufKonsonan != "o" && hurufKonsonan != "u" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}