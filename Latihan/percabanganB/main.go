package main

import "fmt"

func main() {
	fmt.Println("Nama : Muamar Zidan Tri Antoro")
    fmt.Println("NIM : 103012300381")
    fmt.Println("Kelas : IF-47-11")
    fmt.Println("Jawaban soal ppt Percabangan B")
	
	//latihan 1 KTP
	// var cekUmur int
	// var cekKartuKeluarga bool
	// fmt.Scan(&cekUmur, &cekKartuKeluarga)

	// if cekUmur >= 17 && cekKartuKeluarga == true {
	// 	fmt.Println("bisa membuat KTP")
	// } else {
	// 	fmt.Println("belum bisa membuat KTP")
	// }
	//make pseudocode
	//program KTP
	//kamus
	//cekUmur : integer
	//cekKartuKeluarga : boolean
	//algoritma
	//	input(cekUmur, cekKartuKeluarga)
	//	if cekUmur >= 17 and cekKartuKeluarga == true then
	//		output("bisa membuat KTP")
	//	else
	//		output("belum bisa membuat KTP")
	//	endif
	//endprogram

	// latihan 2 Pengiriman Parcek
	// var parcel, parcelKg, parcelGr, biayaTambahan, biaya int;
	// fmt.Scanln(&parcel)

	// parcelKg = parcel / 1000
	// parcelGr = parcel % 1000

	// if parcelKg < 10 && parcelGr >= 500 {
	// 	biayaTambahan = parcelGr * 5
	// 	biaya = (parcelKg * 10000) + biayaTambahan
	// } else if parcel < 10 && parcelGr < 500 {
	// 	biayaTambahan = parcelGr * 15
	// 	biaya = (parcelKg * 10000) + biayaTambahan
	// } else if parcel > 10 {
	// 	biaya  = parcelKg * 10000
	// }
	// fmt.Print("Rp.", parcelKg*10000, " + Rp.", biayaTambahan, " = Rp.", biaya)
	//make pseudocode
	//program Pengiriman_Parcel
	//kamus
	//parcel, parcelKg, parcelGr, biayaTambahan, biaya : integer
	//algoritma
	//	input(parcel)
	//	parcelKg <- parcel / 1000
	//	parcelGr <- parcel % 1000
	//	if parcelKg < 10 and parcelGr >= 500 then
	//		biayaTambahan <- parcelGr * 5
	//		biaya <- (parcelKg * 10000) + biayaTambahan
	//	else if parcel < 10 and parcelGr < 500 then
	//		biayaTambahan <- parcelGr * 15
	//		biaya <- (parcelKg * 10000) + biayaTambahan
	//	else if parcel > 10 then
	//		biaya  <- parcelKg * 10000
	//	endif
	//	output("Rp.", parcelKg*10000, " + Rp.", biayaTambahan, " = Rp.", biaya)
	//endprogram

	//latihan 3 Tiga Bilangan
	// var a, b, c int
	// fmt.Scan(&a, &b, &c)

	// if a > b {
	// 	a, b = b, a
	// }
	// if a > c {
	// 	a, c = c, a
	// }
	// if b > c {
	// 	b, c = c, b
	// }
	// fmt.Println(a, b, c)
	//make pseudocode Tiga_Bilangan
	//program UrutAngka
	//kamus
	//a, b, c : integer
	//algoritma
	//	input(a, b, c)
	//	if a > b then
	//		a, b <- b, a
	//	endif
	//	if a > c then
	//		a, c <- c, a
	//	endif
	//	if b > c then
	//		b, c <- c, b
	//	endif
	//	output(a, b, c)
	//endprogram

	//latihan 4
	var main1, main2, main3, main4, main5 string
	var kalah int
	fmt.Scan(&main1, &main2, &main3, &main4, &main5)

	kalah = 0
	if main1 == "kalah" || main1 == "Kalah" {
		kalah++
	}
	if main2 == "kalah" || main1 == "Kalah" {
		kalah++
	}
	if main3 == "kalah" || main1 == "Kalah" {
		kalah++
	}
	if main4 == "kalah" || main1 == "Kalah" {
		kalah++
	}
	if main5 == "kalah" || main1 == "Kalah" {
		kalah++
	}

	if kalah >= 5 {
		fmt.Println("dipecat")
	} else {
		fmt.Println("tidak dipecat")
	}
	//make pseudocode
	//program Manager_EPL
	//kamus
	//main1, main2, main3, main4, main5 : string
	//kalah : integer
	//algoritma
	//	input(main1, main2, main3, main4, main5)
	//	kalah <-  0
	//	if main1 == "kalah" or main1 == "Kalah" then
	//		kalah <- kalah + 1
	//	endif
	//	if main2 == "kalah" or main1 == "Kalah" then
	//		kalah <- kalah + 1
	//	endif
	//	if main3 == "kalah" or main1 == "Kalah" then
	//		kalah <- kalah + 1
	//	endif
	//	if main4 == "kalah" or main1 == "Kalah" then
	//		kalah <- kalah + 1
	//	endif
	//	if main5 == "kalah" or main1 == "Kalah" then
	//		kalah <- kalah + 1
	//	endif
	//	if kalah >= 5 then
	//		output("dipecat")
	//	else
	//		output("tidak dipecat")
	//	endif
	//endprogram

	//latihan 5
	// var jabatan string;
	// var masaKerja, gaji, anak, tunjanganStaff, tunjanganManager, tunjanganDirektur int;

	// tunjanganStaff = 100;
	// tunjanganManager = 300;
	// tunjanganDirektur = 500;

	// fmt.Scan(&jabatan, &masaKerja, &anak)

	// if jabatan == "staff" && (masaKerja < 5 || masaKerja <= 10)  && anak <= 3 {
	// 	gaji = 4000
	// 	totalGaji := gaji + anak * tunjanganStaff
	// 	fmt.Println(totalGaji)
	// } else if jabatan == "staff" && (masaKerja < 5 || masaKerja <= 10)  && anak > 3 {
	// 	gaji = 4000
	// 	fmt.Println(gaji)
	// } else if jabatan == "manager" && masaKerja > 10 && anak <= 3 {
	// 	gaji = 10000
	// 	totalGaji := gaji + anak * tunjanganManager
	// 	fmt.Println(totalGaji)
	// } else if jabatan == "manager" && masaKerja < 10 && anak < 	3 {
	// 	gaji = 8500
	// 	fmt.Println(gaji)
	// } else if jabatan == "direktur" && masaKerja >= 0 && anak <= 3 {
	// 	gaji = 20000
	// 	totalGaji := gaji + anak * tunjanganDirektur
	// 	fmt.Println(totalGaji)
	// } else {
	// 	fmt.Println("tidak ada")
	// }

	//latihan 9
	// var waktuJam, waktuMenit int
	// var jarak, tarif, totalTarif float64

	// fmt.Scan(&waktuJam, &waktuMenit, &jarak)

	// if waktuJam >= 22 || waktuJam < 5 && waktuMenit > 0 {
	// 	fmt.Println("Maaf, kami tidak bisa melayani pesanan anda.")
	// } else if (waktuJam >= 5 && waktuJam <= 9 || waktuJam >= 16 && waktuJam <= 19) && waktuMenit >= 0 && (jarak >= 0 && jarak <= 10) {
	// 	tarif = 5000.0
	// 	totalTarif = tarif * jarak
	// 	fmt.Println("Total tarif adalah : ", totalTarif)
	// } else if (waktuJam >= 5 && waktuJam <= 9 || waktuJam >= 16 && waktuJam <= 19) && waktuMenit >= 0 && (jarak > 10 && jarak <= 20) {
	// 	tarif = 4500.0
	// 	totalTarif = tarif * jarak
	// 	fmt.Println("Total tarif adalah 	: ", totalTarif)
	// } else {
	// 	tarif = 4000.0
	// 	totalTarif = tarif * jarak
	// 	fmt.Println(totalTarif)
	// }
}
