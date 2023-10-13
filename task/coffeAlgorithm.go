package main

import "fmt"

/*
diketahui :
volume kopi = 500gr
gula = 1000gr
volume 1 sendok = 25ml
massa jenis kopi 0,0025 gr/ml
massa jenis gula 0,001 gr
perbandingan pembuatan kopi = 5 : 1 = gula : kopi agar sama-sama habis
*/

/*
progam coffeAlgorithm

kamus
volKopi : integer = 500
volgula : integer = 1000
kopi : integer = 0

algoritma

for volKopi > 0 and volGula > 0 do
repeat
	output("Ambil gelas")
	output("Masukan 25 ml kopi")
	volKopi - (25 x 0.0025)
	output("Masukan 125 ml gula")
	volGula - (125 x 0.001)
	output("Aduk kopi dan gula")
	output("Kopi ke kopi+1 siap", kopi+1)
	kopi++
until volKopi <= 0 or volGula <= 0

output("Total kopi yang dibuat:", kopi)
if volKopi <= 0 then
	output("Kopi habis. Sisa gula:", volGula, "gr")
end if
if volGula <= 0 then
	output("Gula habis. Sisa kopi:", volKopi, "gr")
end if

endprogram
*/

func main() {
	var volKopi float64 = 500
	var volGula float64 = 1000
	var kopi int = 0
	var volSendokGula float64 = 125
	var volSendokKopi float64 = 25

	for volKopi > 0 && volGula > 0 {
		fmt.Println("Ambil gelas")

		fmt.Println("Masukan 25 ml kopi")
		volKopi = volKopi - (volSendokKopi * 0.0025) //menjadi 0,0625 gr
		fmt.Println("Masukan 125 ml gula")
		volGula = volGula - (volSendokGula * 0.001) //menjadi 0,0625 gr

		fmt.Printf("Kopi ke %d siap\n", kopi+1)
		kopi++
	}

	fmt.Println("Total kopi yang dibuat:", kopi)
	if volKopi <= 0 {
		fmt.Println("Gula habis. Sisa Gula:", volGula, "gr")
	}
	if volGula <= 0 {
		fmt.Println("Kopi habis. Sisa kopi:", volKopi, "gr")
	}
}

/*
diketahui :
volume kopi = 500gr
gula = 1000gr
volume 1 sendok = 25ml
massa jenis kopi 0,0025 gr/ml
massa jenis gula 0,001 gr/ml
perbandingan pembuatan kopi = 2 : 1 = gula : kopi agar salah satu bahan habis duluan
*/

// func main() {
// 	var volKopi float64 = 500
// 	var volGula float64 = 1000
// 	var kopi int = 0

// 	for volKopi > 0 && volGula > 0 {
// 		fmt.Println("Ambil gelas")

// 		fmt.Println("Masukan 25 ml kopi")
// 		volKopi -= 25 * 0.0025 //menjadi 0,0625 gr
// 		fmt.Println("Masukan 50 ml gula")
// 		volGula -= 50 * 0.001 //menjadi 0,05 gr

// 		fmt.Println("Aduk kopi dan gula")
// 		fmt.Println("Kopi ke %d siap", kopi+1)

// 		kopi += 1.0
// 	}

// 	fmt.Println("Total kopi yang dibuat:", kopi)
// 	if volKopi <= 0 {
// 		fmt.Println("Kopi habis. Sisa gula:", volGula, "gr")
// 	}
// 	if volGula <= 0 {
// 		fmt.Println("Gula habis. Sisa kopi:", volKopi, "gr")
// 	}
// }
