package main 

import "fmt"

func main() {
    var n, m, hasil int

    fmt.Scan(&n, m)
    hasil = 0

    for i := 0; i < n; i++ {
        hasil = n + m
    }

    fmt.Println(hasil)
}


