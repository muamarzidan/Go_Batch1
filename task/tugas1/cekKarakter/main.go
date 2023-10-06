package main

import (
	"fmt"
)

//soal no 15
func main() {
    var karakter byte
    fmt.Print("Masukkan karakter: ")
	fmt.Scanf("%c", &karakter)

    fmt.Printf("Karakter yang dimasukkan adalah: %c\n", karakter)
}