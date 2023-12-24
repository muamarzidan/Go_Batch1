package main  

import "fmt"

func main () {
	//latihan1 CetakString 
	// var n, i int;	
	// var text string;

	// fmt.Scan(&n, &text)

	// for i = 0; i < n; i++ {
	// 	fmt.Println(text)
	// }
	//make pseudocode
	// program CetakString
	// kamus
	// 	n, i: integer
	// 	text: string
	// algoritma
	// 	input(n, text)
	// 	for i <- 0 to n do
	// 		output(text)
	// 	endfor
	// endprogram

	//Latihan2 BidangPersegi
	// var n, i, luas, keliling, sisi int
    // fmt.Scan(&n)

    // for i = 0; i < n; i++ {
    //     fmt.Scan(&sisi)

    //     luas = sisi * sisi
    //     keliling = 4 * sisi

    //     fmt.Printf("%d %d\n", luas, keliling)
    // }
	//make pseudocode
	// program BidangPersegi
	// kamus
	// 	n, i, luas, keliling, sisi: integer
	// algoritma
	// 	input(n)
	// 	for i <- 0 to n do
	// 		input(sisi)
	// 		luas <- sisi * sisi
	// 		keliling <- 4 * sisi
	// 		output(luas, keliling)
	// 	endfor
	// endprogram

	//Soal3 Ngoding
	// var i, nMahasiswa, jumlahWaktu, waktuHari int
	// fmt.Scan(&nMahasiswa)

	// for i = 0; i < nMahasiswa; i++ {
	// 	fmt.Scan(&waktuHari)
	// 	jumlahWaktu += waktuHari
	// }

	// rata_rata := float64(jumlahWaktu) / float64(nMahasiswa)
	// fmt.Printf("%.3f\n", rata_rata)
	//make pseudocode
	// program Ngoding
	// kamus
	// 	i, nMahasiswa, jumlahWaktu, waktuHari: integer
	// 	rata_rata: real
	// algoritma
	// 	input(nMahasiswa)
	// 	for i <- 0 to nMahasiswa do
	// 		input(waktuHari)
	// 		jumlahWaktu <- jumlahWaktu + waktuHari
	// 	endfor
	// 	rata_rata <- jumlahWaktu / nMahasiswa
	// 	output(rata_rata)
	// endprogram

	//Soal4 Faktorial
	// var n, i, faktorial int
	// fmt.Scan(&n)

	// faktorial = 1
	// for i = 1; i <= n; i++ {
	// 	faktorial *= i
	// }

	// fmt.Println(faktorial)
	//make pseudocode
	// program Faktorial
	// kamus
	// 	n, faktorial, i: integer
	// algoritma
	// 	input(n)
	// 	faktorial <- 1
	// 	for i <- 1 to n do
	// 		faktorial <- faktorial * i
	// 	endfor
	// 	output(faktorial)
	// endprogram

	//Soal5 FaktorBilangan
	// var n, i int
	// var cekFaktor bool

    // fmt.Scan(&n)
	
    // for i = 1; i <= n; i++ {
	// 	cekFaktor = n % i == 0

	// 	fmt.Println(i , cekFaktor)
	// }
	//make pseudocode
	// program FaktorBilangan
	// kamus
	// 	n, i: integer
	// 	cekFaktor: boolean
	// algoritma
	// 	input(n)
	// 	for i <- 1 to n do
	// 		cekFaktor <- n mod i = 0
	// 		output(i, cekFaktor)
	// 	endfor
	// endprogram
}