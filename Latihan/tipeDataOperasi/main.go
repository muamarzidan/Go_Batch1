package main

import "fmt"

func main() {
	//soal1 MatematikaA

    // var x, y, f float64
    // fmt.Scanf("%f", &x)
    // fmt.Scanf("%f", &y)

    // f = 1 / (3*x*x + 10) + 10*y + 7

    // fmt.Println("f(x, y) = ", f)
    // make pseudo code
    // program MatematikaA
    // kamus
    // x, y, f : float
    // algoritma
    // input(x, y)
    // f <- 1 / (3*x*x + 10) + 10*y + 7
    // output(f)
    // endprogram

    //soal2 MatematikaB
    // var x, f float64
    // fmt.Scanf("%f\n", &x)

    // f = ((x*x*x) + (3*x)) / ((x*x*x*x) - 3*x*x + 4)

    // fmt.Printf("f(x) %f\n", f)
    //make pseudo code
    // program MatematikaB
    // kamus
    // x, f : float
    // algoritma
    // input(x)
    // f <- ((x*x*x) + (3*x)) / ((x*x*x*x) - 3*x*x + 4)
    // output(f)
    // endprogram

    //soal3 TigaDigit
    var x, d1, d2, d3 int
    fmt.Scan(&x)

    d1 = x / 100
    d2 = (x % 100) / 10
    d3 = x % 10

    fmt.Print(d1, d2, d3)
    //make pseudo code
    program TigaDigit
    kamus
    x, d1, d2, d3 : int
    algoritma
    input(x)
    d1 <- x / 100
    d2 <- (x % 100) / 10
    d3 <- x % 10
    output(d1, d2, d3)
    endprogram
}