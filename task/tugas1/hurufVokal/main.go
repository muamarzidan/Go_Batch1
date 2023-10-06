package main

import (
	"fmt"
	"strings"
)

func main() {
	var karakter string

	fmt.Print("Masukkan huruf : ")
	fmt.Scan(&karakter)
	karakter = strings.ToLower(karakter)

	if karakter == "a" || karakter == "e" || karakter == "i" || karakter == "o" || karakter == "u" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}