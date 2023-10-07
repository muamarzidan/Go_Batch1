package main

import (
	"fmt"
)


//soal no 10
func main() {
    var bilNegatif int
    var Hasil bool

    fmt.Print("Masukkan bilangan bulat : ")
    fmt.Scan(&bilNegatif)

    if bilNegatif < 0 {
        Hasil = true
        fmt.Println(Hasil)
    } else {
        Hasil = false
        fmt.Println(Hasil)
    }
}

//soal no 11
// func main() {
//     var bilPositif int
// 	var Hasil bool

//     fmt.Print("Masukkan bilangan bulat : ")
//     fmt.Scan(&bilPositif)

//     if bilPositif > 0 {
// 		Hasil = true
// 		fmt.Println(Hasil)
//     } else {
// 		Hasil = false
// 		fmt.Println(Hasil)
//     }
// }












