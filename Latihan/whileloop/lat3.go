package main

import "fmt"

func main() {
    var x, xi, y, i int
    fmt.Scan(&x, &y)
    i = 0
    xi = x

    for xi >= y {
        xi = xi - y
        i = i + 1
    }

    fmt.Println(x, "mod", y, "=", xi)
    fmt.Println(x, "div", y, "=", i)
}
