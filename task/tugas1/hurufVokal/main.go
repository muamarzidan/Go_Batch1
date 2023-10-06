package main

import (
	"fmt"
	"strings"
)


//soal no 12
func main() {
	var hurufVokal string

	fmt.Print("Masukkan huruf : ")
	fmt.Scan(&hurufVokal)
	hurufVokal = strings.ToLower(hurufVokal)

	if hurufVokal == "a" || hurufVokal == "e" || hurufVokal == "i" || hurufVokal == "o" || hurufVokal == "u" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}