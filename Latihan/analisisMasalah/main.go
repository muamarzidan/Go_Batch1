package main 

import "fmt"

func main() {
	//cek bilangan prima
	// var bilangan int
	// fmt.Print("Masukkan bilangan : ")
	// fmt.Scan(&bilangan)

	// var cek bool
	// cek = true

	// if bilangan == 1 || bilangan == 0 {
	// 	cek = false
	// 	fmt.Println(cek)
	// } else if bilangan == 2 {
	// 	fmt.Println(cek)
	// }  else if (bilangan / 1 == bilangan && bilangan / bilangan == 1) && (bilangan % 2 != 0){
	// 	fmt.Println(cek)
	// } else {
	// 	cek = false
	// 	fmt.Println(cek)
	// }

	//Latihan1 Faktor Bilangan
	// var x, y int
	// var cekFaktor bool

	// cekFaktor = true
	// fmt.Scan(&x, &y)

	// cekFaktor = y % x == 0
	// fmt.Println(cekFaktor)
	//make pseudocode
	// program FaktorBilangan
	// kamus
	// x, y : integer
	// algoritma
	// input (x, y)
	// cekFaktor <- y % x == 0
	// output(cekFaktor)
	// endprogram

	//Latihan2 Penduduk
	// var totalPenduduk, penduduk, kelahiran, imigrasi, kematian, emigrasi int
	// fmt.Scan(&penduduk, &kelahiran, &imigrasi, &kematian, &emigrasi)

	// totalPenduduk = penduduk + kelahiran + imigrasi - kematian - emigrasi
	// fmt.Println(totalPenduduk)
	//make pseudocode
	// program Penduduk
	// kamus
	// totalPenduduk, penduduk, kelahiran, imigrasi, kematian, emigrasi : integer
	// algoritma
	// input(penduduk, kelahiran, imigrasi, kematian, emigrasi)
	// totalPenduduk <- penduduk + kelahiran + imigrasi - kematian - emigrasi
	// output(totalPenduduk)
	// endprogram

	//latihan3	Cek Bilangan Akar Pangkat 3
	// var x, y int
	// var cek bool
	// cek = true

	// fmt.Scan(&x, &y)
	// cek = y == x * x * x
	// fmt.Println(cek)
	//make pseudocode
	// program CekBilanganAkarPangkat3
	// kamus
	// x, y : integer
	// cek : boolean
	// algoritma
	// input(x, y)
	// cek <- y == x * x * x
	// output(cek)
	// endprogram

	//latihan4 Gravitasi
	// var merkurius, venus, mars, bumi, yupiter, saturnus, uranus, neptunus, x float64
	// bumi = 9.8

	// merkurius = 0.38 * bumi
	// venus = 0.91 * bumi
	// mars = 0.38 * bumi
	// yupiter = 2.37 * bumi
	// saturnus = 0.92 * bumi
	// uranus = 0.89 * bumi
	// neptunus = 1.13 * bumi

	// fmt.Scan(&x)

	// fmt.Printf("%.0f %.0f %.0f %.0f %.0f %.0f %.0f %.0f\n", x * merkurius, x * venus, x * bumi, x * mars, x * yupiter, x * saturnus, x * uranus, x * neptunus)
	//make pseudocode
	// program Gravitasi
	// kamus
	// merkurius, venus, mars, bumi, yupiter, saturnus, uranus, neptunus, x : real
	// algoritma
	//  input(x)
	//  merkurius <- 0.38 * bumi
	//  venus <- 0.91 * bumi
	//  mars <- 0.38 * bumi
	//  yupiter <- 2.37 * bumi
	//  saturnus <- 0.92 * bumi
	//  uranus <- 0.89 * bumi
	//  neptunus <- 1.13 * bumi
	//  utput(x * merkurius, x * venus, x * bumi, x * mars, x * yupiter, x * saturnus, x * uranus, x * neptunus)
	// endprogram

	//latihan5
	// var suku1, suku2, suku3, suku4, suku5 int
	// fmt.Scan(&suku1, &suku2)
	// suku3 = suku2 + suku1
	// suku4 = suku3 + suku2
	// suku5 = suku4 + suku3
	// fmt.Println(suku3, suku4, suku5)
	//make pseudocode
	// program DeretAritmatika
	// kamus
	//     suku1, suku2, suku3, suku4, suku5 : integer
	// algoritma
	// 	input(suku1, suku2)
	// 	suku3 <- suku2 + suku1
	// 	suku4 <- suku3 + suku2
	// 	suku5 <- suku4 + suku3
	// 	output(suku3, suku4, suku5)
	// endprogram
}

