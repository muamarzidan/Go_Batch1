package main

import "fmt"

func main() {
	//Soal1 Waktu
	// var t, jt, mt, dt int
	// fmt.Scan(&t)

	// jt = t / 3600
	// mt = (t % 3600) / 60
	// dt = (t % 3600) % 60

	// fmt.Println(jt, "jam", mt, "menit", dt, "detik")
	//make pseudocode
	//program Waktu
	//kamus
	//t, jt, mt, dt : integer
	//algoritma
	//input(t)
	//jt <- t / 3600
	//mt <- (t % 3600) / 60
	//dt <- (t % 3600) % 60
	//output(jt, "jam", mt, "menit", dt, "detik")
	//endprogram

	//Soal2 Voucher
	// var x, d1, d3, d2d3, d4 int
	// var diskon, cashback bool
	// diskon = true
	// cashback = true

	// fmt.Scan(&x)

	// d1 = x / 1000
	// d3 = (x % 100) / 10
	// d4 = x % 10
	// d2d3 = (x % 100) / 10

	// diskon = d2d3%2 == 0

	// cashback = (d4 != 0) && ((d1 + d3) % d4 == 0)

	// fmt.Println("Diskon?", diskon, "Cashback?", cashback)
	//make pseudocode
	//program Voucher
	//kamus
	//x, d1, d3, d2d3, d4 : integer
	//diskon, cashback : boolean
	//algoritma
	//input(x)
	//d1 <- x / 1000
	//d3 <- (x % 100) / 10
	//d4 <- x % 10
	//d2d3 <- (x % 100) / 10
	//diskon <- d2d3%2 == 0
	//cashback <- (d4 != 0) && ((d1 + d3) % d4 == 0)
	//output("Diskon?", diskon, "Cashback?", cashback)
	//endprogram	

	//Soal3 Jumlah Bilangan
	// var x, x2, total int
    // fmt.Scan(&x)

    // for x != 0 {
    //     x2 = x % 10
    //     x = x / 10
    //     total = total + x2
    // }
	// fmt.Println(total)
	//make pseudocode
	//program Jumlah_Bilangan
	//kamus
	//x, x2, total : integer
	//algoritma
	//input(x)
	//while x != 0 do
	//	x2 <- x % 10
	//	x <- x / 10
	//	total <- total + x2
	//end while
	//output(total)
	//endprogram

	//Soal4 Konsekutif
	// var X, angka, digitSebelum, selisih int
	// var cekKonsekutif bool
	// fmt.Scan(&X)

	// digitSebelum = -1
	// cekKonsekutif = true

	// for X > 0 {
	// 	angka = X % 10

	// 	if digitSebelum != -1 {
	// 		selisih = digitSebelum - angka
	// 		if selisih != 1 && selisih != -1 {
	// 			cekKonsekutif = false
	// 			break
	// 		}
	// 	}

	// 	digitSebelum = angka
	// 	X /= 10
	// }

	// fmt.Println(cekKonsekutif)
	//make pseudocode
	//program Konsekutif
	//kamus
	// X, angka, digitSebelum, selisih : integer
	// cekKonsekutif : boolean
	//algoritma
	// input(X)
	// digitSebelum <- -1
	// cekKonsekutif <- true
	// while X > 0 do
	// 	angka <- X mod 10
	// 	if digitSebelum != -1 then
	// 		selisih <- digitSebelum - angka
	// 		if selisih != 1 and selisih != -1 then
	// 			cekKonsekutif <- false
	// 			break
	// 		endif
	// 	endif
	// 	digitSebelum <- angka
	// 	X <- X div 10
	// endwhile
	// output(cekKonsekutif)
	//endprogram

	//Soal5 Fibonacci
	var n, fibo0, fibo1 int
    fmt.Scan(&n)

    fibo0 = 0
    fibo1 = 1

    fmt.Print(fibo0)

    for i := 1; i <= n; i++ {
        fmt.Print(" ", fibo1)
        fibo0, fibo1 = fibo1, fibo0+fibo1
    }
    fmt.Println()
	//make pseudocode
	//program Fibonacci
	//kamus
	//n, fibo0, fibo1 : integer
	//algoritma
	//input(n)
	//fibo0 <- 0
	//fibo1 <- 1
	//output(fibo0)
	//for i <- 1 to n do
	//	output(fibo1)
	//	fibo0 <- fibo1
	//	fibo1 <- fibo0 + fibo1
	//endfor
	//output()
	//endprogram
}