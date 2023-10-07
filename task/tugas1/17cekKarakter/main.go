package main

import (
	"fmt"
)

//soal no 17
func main() {
    var karakter byte
    fmt.Print("Masukkan sebuah karakter : ")
	fmt.Scanf("%c", &karakter)

    fmt.Printf("Karakter yang dimasukkan adalah : %c\n", karakter)
}