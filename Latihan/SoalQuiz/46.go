package main

import "fmt"

func main() {
    var X, X2 int;
    fmt.Scan(&X);

    for X != 0 {
        X2 = X % 10
        X = X / 10
        fmt.Print(X2)
    }
}

// program balikangka
// kamus
//     X, X2 : integer
// algoritma
//     input(X)
//     while X != 0 do 
//         X2 <- X mod 10
//         X <- X div 10
//         output(X2)
//     end while
// end program



