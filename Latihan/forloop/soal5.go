package main

import "fmt"

func main() {
    var n, i int
    fmt.Scan(&n)
    
    for i = 1; i <= n; i++ {
        fmt.Printf("%d %t\n", i, n % i == 0)
    }
}