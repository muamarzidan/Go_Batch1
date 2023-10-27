package main

import "fmt"

func main() {
    var n, i int
    fmt.Scan(&n)
	
    for i = 1; i <= n; i++ {
        if n % i == 0 {
            fmt.Printf("%d true\n", i)
        } else {
            fmt.Printf("%d false\n", i)
        }
    }
}