package main

import "fmt"

func main() {
    var x, total int;
    total = 0;
    for x%2 == 0{
        fmt.Scan(&x)
        total = total + x
    }
    fmt.Println(total);
}
