package main

import "fmt"

func main() {
	// var konsonan string;
	// fmt.Scan(&konsonan);

	// if (konsonan == "a" || konsonan == "i" || konsonan == "u" || konsonan == "e" || konsonan == "o") || (konsonan == "A" || konsonan == "I" || konsonan == "U" || konsonan == "E" || konsonan == "O"){
	// 	fmt.Println("bukan konsonan");
	// } else {
	// 	fmt.Println("konsonan");
	// }
	

	var angka int
    fmt.Scan(&angka)

    if angka%3 == 0 && angka%5 == 0 {
        fmt.Println("Kelipatan 3 dan Kelipatan 5")
    } else if angka%3 == 0 {
        fmt.Println("Kelipatan 3")
    } else if angka%5 == 0 {
        fmt.Println("Kelipatan 5")
    }
}