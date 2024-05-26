package main

import "fmt"

func main() {
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

	// buatlah programuntuk menghitung jumlah deret :4/n + 4/n+1 + 4/n+2 + ... + 4/m
	// dengan n dan m adalah bilangan positif dan n < m
	// masukan terdiri dari dua bilangan bulat positif n dan m
	// keluaran berupa jumlah deret tersebut dengan pembulatan 2 angka di belakang koma
	// contoh masukan = 1 dan 5
	// contoh keluaran = 9.13

	var n, m int
	var jumlahDeret float64
	fmt.Scan(&n, &m)

	for i := n; i <= m; i++ {
		jumlahDeret += 4 / float64(i)
	}
	fmt.Printf("%.2f\n", jumlahDeret)

}

// program cekKonsekutif

// kamus
// x, angka, digitSebelum, selisih : integer

// algoritma
// input(x)
// digitSebelum <- -1
// cekKonsekutif <- true
// while x > 0
// 	angka <- X mod 10
// 	if digitSebelum not -1
// 		selisih <- digitSebelum - angka
// 		if selisih not 1 and selisih not -1
// 			cekKonsekutif <- false
// 			break
// 	digitSebelum <- angka
// 	X <- x div 10
// print(cekKonsekutif)

// endprogram

// var total, jumlahKali, input int

// for total <= 100 {
// 	fmt.Print("Masukkan angka: ")
// 	fmt.Scanln(&input)
// 	total += input
// 	jumlahKali++
// }
