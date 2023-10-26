package main

import "fmt"

func main() {
    var x, total int
    total = 0
    for {
        fmt.Scan(&x)
        if x%2 == 0 {
            total = total + x
        } else {
            break
        }
    }
    fmt.Println(total)
}
