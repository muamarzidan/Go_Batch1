package main

import "fmt"

func main() {
    var nGula, mKopi, xGula, yKopi, cangkir int;
    fmt.Scan(&nGula, &mKopi, &xGula, &yKopi);

    xGula = nGula;
    yKopi = mKopi;
    cangkir = 0;

    for xGula <= nGula && yKopi <= mKopi {
        xGula = xGula - 1;
        yKopi = yKopi - 1;
        cangkir = cangkir + 1;
    }

    fmt.Println(cangkir);
}
