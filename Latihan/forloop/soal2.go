package main

import "fmt"

func main() {
    var n, i, sisi int
    fmt.Scan(&n)

    for i = 0; i < n; i++ {
        fmt.Scan(&sisi)

        luas := sisi * sisi
        keliling := 4 * sisi

        fmt.Println(luas, keliling)
    }
}
