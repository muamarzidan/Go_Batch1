package main

import "fmt"

func main() {
    var X, X2, total int
    fmt.Scan(&X)

    for X != 0 {
        X2 = X % 10
        X = X / 10
        fmt.Print(X2)
        total = total + X2
    }

    fmt.Println()
    fmt.Println(total)
}
