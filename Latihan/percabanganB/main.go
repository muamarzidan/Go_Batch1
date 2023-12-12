package main 

import "fmt"

func main () {
	//percabangan B

	//latihan 1
	// var cekUmur int;
	// var cekKartuKeluarga bool;

	// fmt.Scan(&cekUmur, &cekKartuKeluarga);

	// if cekUmur >= 17 && cekKartuKeluarga == true {
	// 	fmt.Println("bisa membuat KTP");
	// } else {
	// 	fmt.Println("belum bisa membuat KTP");
	// }


	// latihan 2
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


	//latihan 3
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

	
	//latihan 4
	// var match1, match2, match3, match4, match5 string
	// var kalah int;
	// fmt.Scan(&match1, &match2, &match3, &match4, &match5)

	// kalah =  0
	// if match1 == "kalah" {
	// 	kalah++
	// }
	// if match2 == "kalah" {
	// 	kalah++
	// }
	// if match3 == "kalah" {
	// 	kalah++
	// }
	// if match4 == "kalah" {
	// 	kalah++
	// }
	// if match5 == "kalah" {
	// 	kalah++
	// }

	// if kalah >= 5 {
	// 	fmt.Println("dipecat")
	// } else {
	// 	fmt.Println("tidak dipecat")
	// }


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
	var waktuJam, waktuMenit int;
	var jarak, tarif, totalTarif float64;

	fmt.Scan(&waktuJam, &waktuMenit, &jarak)

	if waktuJam >= 22 || waktuJam < 5 && waktuMenit > 0 {
		fmt.Println("Maaf, kami tidak bisa melayani pesanan anda.")
	} else if (waktuJam >= 5 && waktuJam <= 9 || waktuJam >=16 && waktuJam <=19) && waktuMenit >= 0 && (jarak >= 0 && jarak <= 10) {
		tarif = 5000.0
		totalTarif = tarif * jarak 
		fmt.Println("Total tarif adalah : ", totalTarif)
	} else if (waktuJam >= 5 && waktuJam <= 9 || waktuJam >=16 && waktuJam <=19) && waktuMenit >= 0 && (jarak > 10 && jarak <= 20) {
		tarif = 4500.0
		totalTarif = tarif * jarak
		fmt.Println("Total tarif adalah 	: ", totalTarif)
	} else {
		tarif = 4000.0
		totalTarif = tarif * jarak
		fmt.Println(totalTarif)
	}
}