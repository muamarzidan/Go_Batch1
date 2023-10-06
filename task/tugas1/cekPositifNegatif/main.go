package main

import (
	"fmt"
	// "strings"
)

//9
// func main() {
//     var bilangan int

//     fmt.Print("Masukkan bilangan bulat: ")
//     fmt.Scan(&bilangan)

//     if bilangan < 0 {
//         fmt.Println("true")
//     } else {
//         fmt.Println("false")
//     }
// }

//10
func main() {
    var bilangan int

    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)

    if bilangan < 0 {
        fmt.Println("false")
    } else {
        fmt.Println("true")
    }
}





// 13
// func main() {
// 	var bilangan int

// 	fmt.Print("Masukkan sebuah bilangan bulat: ")
// 	fmt.Scan(&bilangan)

// 	if bilangan == 0 {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}
// }

// 14
// func main() {
// 	var nama string

// 	fmt.Print("Masukkan kata : ")
// 	fmt.Scan(&nama)

// 	fmt.Println("Hai " + nama + ", saya sedang belajar Golang")
// }

//15
// func main() {
//     var karakter byte
//     fmt.Print("Masukkan karakter: ")
// 	fmt.Scanf("%c", &karakter)


//     fmt.Printf("Karakter yang dimasukkan adalah: %c\n", karakter)
// }


