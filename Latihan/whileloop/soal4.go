package main

import "fmt"

func main() {
    var nGula, mKopi, xGula, yKopi, cangkir int
    fmt.Scan(&nGula, &mKopi, &xGula, &yKopi)

    for xGula <= nGula && yKopi <= mKopi {
        nGula -= xGula
        mKopi -= yKopi
        cangkir++
    }

    fmt.Println(cangkir)
}
