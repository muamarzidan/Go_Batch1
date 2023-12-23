package main

import "fmt"

func main() {
	//Soal1 Login
    // var inputUsername, inputPassword string;
    // var i, login int

	// i = 1
	// login = -1

    // for i > 0 && (inputUsername != "admin" || inputPassword != "admin") {
    //     fmt.Scan(&inputUsername)
    //     fmt.Scan(&inputPassword)
    //     login++
    // };

    // fmt.Println(login)
    // fmt.Print("Login berhasil")
	//make pseudocode
	//program Login
	//kamus
	// inputUsername, inputPassword : string
	// i, login : integer
	//algoritma
	// i <- 1
	// login <- -1
	// while i > 0 and (inputUsername != "admin" or inputPassword != "admin") do
	// 	input(inputUsername)
	// 	input(inputPassword)
	// 	login <- login + 1
	// end while
	// output(login)
	// output("Login berhasil")
	//end program

	//Soal2 Dompet
	// var saldo, tf int;
	// saldo = 0;

	// for tf != 0 {
	// 	saldo = saldo + tf
	// 	fmt.Scan(&tf)
	// };
	// fmt.Println(saldo);
	//make pseudocode
	//program Dompet
	//kamus
	// saldo, tf : integer
	//algoritma
	// saldo <- 0
	// while tf != 0 do
	// 	saldo <- saldo + tf
	// 	input(tf)
	// end while
	// output(saldo)
	//end program

	//Soal3 Digit
	// var x, x2, total int
    // fmt.Scan(&x)

    // for x != 0 {
    //     x2 = x % 10
    //     x = x / 10
    //     fmt.Print(x2)
    //     total = total + x2
    // }
	// fmt.Println()
	// fmt.Println(total)
	//make pseudocode
	//program Digit
	//kamus
	// X, X2, total : integer
	//algoritma
	// input(X)
	// while X != 0 do
	// 	X2 <- X mod 10
	// 	X <- X div 10
	// 	output(X2)
	// 	total <- total + X2
	// end while
	// output(total)
	//end program

	//Soal4 Cnagkir Kopi
	// var n, m, x, y, cangkir int
    // fmt.Scan(&n, &m, &x, &y)

    // cangkir = 0
    // for n >= x && m >= y {
    //     n = n - x
    //     m = m - y
    //     cangkir++
    // }
    // fmt.Println(cangkir)
	//make pseudocode
	//program CangkirKopi
	//kamus
	// n, m, x, y, cangkir : integer
	//algoritma
	// input(n, m, x, y)
	// cangkir <- 0
	// while n >= x and m >= y do
	// 	n <- n - x
	// 	m <- m - y
	// 	cangkir <- cangkir + 1
	// end while
	// output(cangkir)
	//end program

	//Soal5 Konsekutif
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

	//Latihan6 Tanki Air
	var t, ember, volume int;
	var cekTangki bool
	ember = 0;
	cekTangki = true
	fmt.Scan(&t);
	
	for t != ember {
		fmt.Scan(&volume)
		ember = ember + volume
		cekTangki = ember == t
		fmt.Println(cekTangki)
	}
	//make pseudocode
	//program TankiAir
	//kamus
	// t, ember, volume : integer
	// cekTangki : boolean
	//algoritma
	// ember <- 0
	// cekTangki <- true
	// input(t)
	// while t != ember do
	// 	input(volume)
	// 	ember <- ember + volume
	// 	cekTangki <- ember == t
	// 	output(cekTangki)
	// endwhile
	//endprogram
}
