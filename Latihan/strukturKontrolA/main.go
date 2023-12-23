package main 

import "fmt"

func main() {
	//Soal1 Faktor Bilangan
	// var x, i int;

	// fmt.Scan(&x)
	// fmt.Println("");

	// for i = 1; i <= x; i++ {
	// 	if x % i == 0 {
	// 		fmt.Print(i, " ")
	// 	}
	// }
	//make pseudocode
	//program Faktor_Bilangan
	//kamus
	//	x, i : integer
	//algoritma
	//	input(x)
	//	for i <- 1 to x do
	//		if x mod i = 0 then
	//			output(i)
	//		endif
	//	endfor
	//endprogram

	//Soal2 Bilangan Prima
	// var x int;
	// var prima bool;

	// fmt.Scan(&x)
	
	// if x == 1 {
	// 	prima = false
	// 	fmt.Print(prima)
	// } else if x == 2 {
	// 	prima = true
	// 	fmt.Print(prima)
	// } else if x / x == 1 && x / 1 == x && x % 2 != 0 {
	// 	prima = true
	// 	fmt.Print(prima)
	// } else {
	// 	prima = false
	// 	fmt.Print(prima)
	// }
	//make pseudocode
	//program Bilangan_Prima
	//kamus
	//	x : integer
	//	prima : boolean
	//algoritma
	//	input(x)
	//	if x = 1 then
	//		prima <- false
	//		output(prima)
	//	else if x = 2 then
	//		prima <- true
	//		output(prima)
	//	else if x / x = 1 and x / 1 = x and x mod 2 != 0 then
	//		prima <- true
	//		output(prima)
	//	else
	//		prima <- false
	//		output(prima)
	//	endif
	//endprogram

	//Soal3 Biner
	var x int
	fmt.Scan(&x)

	var biner string
	biner = ""

	for x > 0 {
		if x % 2 == 0 {
			biner = "0" + biner
		} else {
			biner = "1" + biner
		}
		x = x / 2
	}
	fmt.Println(biner)
	//make pseudocode
	//program Biner
	//kamus
	//	x : integer
	//	biner : string
	//algoritma
	//	input(x)
	//	biner <- ""
	//	while x > 0 do
	//		if x mod 2 = 0 then
	//			biner <- "0" + biner
	//		else
	//			biner <- "1" + biner
	//		endif
	//		x <- x div 2
	//	endwhile
	//	output(biner)
	//endprogram
	
	//Soal4 Lebar Daun
	// var n, lebar, max int
    // fmt.Scan(&n)

    // fmt.Scan(&max)
    // for i := 1; i < n; i++ {
    //     fmt.Scan(&lebar)
    //     if lebar > max {
    //         max = lebar
    //     }
    // }
    // fmt.Println(max)
	//make pseudocode
	//program Lebar_Daun
	//kamus
	//	n, lebar, max : integer
	//algoritma
	//	input(n)
	//	input(max)
	//	for i <- 1 to n do
	//		input(lebar)
	//		if lebar > max then
	//			max <- lebar
	//		endif
	//	endfor
	//	output(max)
	//endprogram

	//Soal5 N Digit
	// var bilBulat, max int
	// fmt.Scan(&bilBulat)

	// max = bilBulat % 10
	// bilBulat = bilBulat / 10

	// for bilBulat > 0 {
	// 	digit := bilBulat % 10
	// 	if digit > max {
	// 		max = digit
	// 	}
	// 	bilBulat /= 10
	// }

	// fmt.Println(max)
	//make pseudocode
	//program N_Digit
	//kamus
	//	bilBulat, max : integer
	//algoritma
	//	input(bilBulat)
	//	max <- bilBulat mod 10
	//	bilBulat <- bilBulat div 10
	//	while bilBulat > 0 do
	//		digit <- bilBulat mod 10
	//		if digit > max then
	//			max <- digit
	//		endif
	//		bilBulat <- bilBulat div 10
	//	endwhile
	//	output(max)
	//endprogram

	//Soal6 Cari Digit
	// var x, n, digitPisah int
    // var cekAda bool
    // fmt.Scan(&x, &n)

    // cekAda = false
    // for n > 0 && !cekAda {
    //     digitPisah = n % 10
    //     if digitPisah == x {
    //         cekAda = true
    //     }
    //     n /= 10
    // }

    // fmt.Println(cekAda)
	//make pseudocode
	//program Cari_Digit
	//kamus
	//	x, n, digitPisah : integer
	//	cekAda : boolean
	//algoritma
	//	input(x, n)
	//	cekAda <- false
	//	while n > 0 and not cekAda do
	//		digitPisah <- n mod 10
	//		if digitPisah = x then
	//			cekAda <- true
	//		endif
	//		n <- n div 10
	//	endwhile
	//	output(cekAda)
	//endprogram

	//Soal7 Gerbang Tol
	// var inputHuruf string
	// var cekInput bool

	// cekInput = true

	// var totalA, totalB, totalC int
	// totalA = 0
	// totalB = 0
	// totalC = 0

	// for cekInput {
	// 	fmt.Scan(&inputHuruf)

	// 	if inputHuruf == "A" {
	// 		totalA++
	// 	} else if inputHuruf == "B" {
	// 		totalB++
	// 	} else if inputHuruf == "C" {
	// 		totalC++
	// 	} else {
	// 		cekInput = false
	// 	}
	// }
	// fmt.Println("Tipe A = ", totalA)
	// fmt.Println("Tipe B = ", totalB)
	// fmt.Println("Tipe C = ", totalC)
	//make pseudocode
	//program Gerbang_Tol
	//kamus
	//	inputHuruf : string
	//	cekInput : boolean
	//	totalA, totalB, totalC : integer
	//algoritma
	//	cekInput <- true
	//	totalA <- 0
	//	totalB <- 0
	//	totalC <- 0
	//	while cekInput do
	//		input(inputHuruf)
	//		if inputHuruf = "A" then
	//			totalA <- totalA + 1
	//		else if inputHuruf = "B" then
	//			totalB <- totalB + 1
	//		else if inputHuruf = "C" then
	//			totalC <- totalC + 1
	//		else
	//			cekInput <- false
	//		endif
	//	endwhile
	//	output("Tipe A = ", totalA)
	//	output("Tipe B = ", totalB)
	//	output("Tipe C = ", totalC)
	//endprogram

	//Soal8 Tempratur
	// var x, min, max, cekX, jumlahX, n int;
	// var rataRata float64;
	// var cekUlang bool;

	// cekUlang = true
	// rataRata, min, max, cekX, jumlahX, n = 0, 0, 0, 0, 0, 0;

	// for cekUlang == true {
	// 	fmt.Scan(&x)
	// 	if cekX == 0 && x == 0 {
	// 		cekUlang = false;
	// 	}
	// 	if min > x {
	// 		min = x
	// 	}
	// 	if max < x {
	// 		max = x
	// 	}

	// 	jumlahX = jumlahX + x
	// 	cekX = x
	// 	n++
	// }

	// rataRata = float64(jumlahX) / float64(n-1)
	
	// fmt.Println("")
	// fmt.Println(max)
	// fmt.Println(min)
	// fmt.Println(rataRata)
	//make pseudocode	
	//program Temperatur
	//kamus 
	//	x, min, max, cekX, jumlahX, n : integer
	//	rataRata : float
	//	cekUlang : boolean
	//algoritma
	//	cekUlang <- true
	//	rataRata, min, max, cekX, jumlahX, n <- 0
	//	while cekUlang = true do
	//		input(x)
	//		if cekX = 0 and x = 0 then
	//			cekUlang <- false
	//		endif
	//		if min > x then
	//			min <- x
	//		endif
	//		if max < x then
	//			max <- x
	//		endif
	//		jumlahX <- jumlahX + x
	//		cekX <- x
	//		n <- n + 1
	//	endwhile
	//	rataRata <- jumlahX / (n - 1)
	//	output(max)
	//	output(min)
	//	output(rataRata)
	//endprogram

	//Soal9 Pola Bilangan 1
	// var x, i, j int
	// fmt.Scan(&x)

	// for i = 1; i <= x; i++ {
	// 	for j = 1; j <= x; j++ {
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println("")
	// }
	//make pseudocode
	//program Pola_Bilangan_1
	//kamus
	//	x, i, j : integer
	//algoritma
	//	input(x)
	//	for i <- 1 to x do
	//		for j <- 1 to x do
	//			output(j, " ")
	//		endfor
	//		output("")
	//	endfor
	//endprogram

	//Soal10 Pola Bilangan 2
	// var x, i, j int
	// fmt.Scan(&x)
	
	// for i = 1; i <= x; i++ {
	// 	for j = 1; j <= x; j++ {
	// 		fmt.Print(i, " ")
	// 	}
	// 	fmt.Println("")
	// }
	//make pseudocode
	//program Pola_Bilangan_2
	//kamus
	//	x, i, j : integer
	//algoritma
	//	input(x)
	//	for i <- 1 to x do
	//		for j <- 1 to x do
	//			output(i, " ")
	//		endfor
	//		output("")
	//	endfor
	//endprogram

	//Soal11 Pola Bilangan 3
	// var x, i, j int
	// fmt.Scan(&x)

	// for i = 1; i <= x; i++ {
	// 	for j = 1; j <= x; j++ {
	// 		if i == 1 || i == x || j == 1 || j == x {
	// 			fmt.Print(i, " ")
	// 		} else {
	// 			fmt.Print("  ")
	// 		}
	// 	}
	// 	fmt.Println("")
	// }
	//make pseudocode
	//program Pola_Bilangan_3
	//kamus
	//	x, i, j : integer
	//algoritma
	//	input(x)
	//	for i <- 1 to x do
	//		for j <- 1 to x do
	//			if i = 1 or i = x or j = 1 or j = x then
	//				output(i, " ")
	//			else
	//				output("  ")
	//			endif
	//		endfor
	//		output("")
	//	endfor
	//endprogram

	//Soal12 Pola Bilangan 4
	// var x, i, j int
	// fmt.Scan(&x)
		
	// for i = 1; i <= x; i++ {
	// 	for j = 1; j <= x; j++ {
	// 		if i == j || i + j == x + 1 {
	// 			fmt.Print(i, " ")
	// 		} else {
	// 			fmt.Print("  ")
	// 		}
	// 	}
	// 	fmt.Println("")
	// }
	//make pseudocode
	//program Pola_Bilangan_4
	//kamus
	//	x, i, j : integer
	//algoritma
	//	input(x)
	//	for i <- 1 to x do
	//		for j <- 1 to x do
	//			if i = j or i + j = x + 1 then
	//				output(i, " ")
	//			else
	//				output("  ")
	//			endif
	//		endfor
	//		output("")
	//	endfor
	//endprogram
}