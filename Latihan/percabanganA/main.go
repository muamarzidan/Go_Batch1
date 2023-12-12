package main

import "fmt"

func main() {
	//latihan1 Konsonan
	// var konsonan string;
	// fmt.Scan(&konsonan);

	// if (konsonan == "a" || konsonan == "i" || konsonan == "u" || konsonan == "e" || konsonan == "o") || (konsonan == "A" || konsonan == "I" || konsonan == "U" || konsonan == "E" || konsonan == "O"){
	// 	fmt.Println("bukan konsonan");
	// } else {
	// 	fmt.Println("konsonan");
	// }
	//make pseudocode
	//program Konsonan
	//kamus
	//konsonan : string
	//algoritma
	//	input(konsonan)
	//	if (konsonan == "a" || konsonan == "i" || konsonan == "u" || konsonan == "e" || konsonan == "o") || (konsonan == "A" || konsonan == "I" || konsonan == "U" || konsonan == "E" || konsonan == "O") then
	//		output("bukan konsonan")
	//	else
	//		output("konsonan")
	//	endif
	//endprogram

	//latihan2 Kelipatan
	// var angka int
	// fmt.Scan(&angka)

	// if angka%3 == 0 && angka%5 == 0 {
	// 	fmt.Println("Termasuk kelipatan 3 dan kelipatan 5")
	// } else if angka%3 == 0 {
	// 	fmt.Println("Kelipatan 3")
	// } else if angka%5 == 0 {
	// 	fmt.Println("Kelipatan 5")
	// }
	//make pseudocode
	//program Kelipatan
	//kamus
	//angka : integer
	//algoritma
	//	input(angka)
	//	if angka%3 == 0 && angka%5 == 0 then
	//		output("Kelipatan 3 dan Kelipatan 5")
	//	else if angka%3 == 0 then
	//		output("Kelipatan 3")
	//	else if angka%5 == 0 then
	//		output("Kelipatan 5")
	//	endif
	//endprogram

	//latihan3 Segitiga
	// var sisiA, sisiB, sisiC int
	// fmt.Scan(&sisiA, &sisiB, &sisiC)

	// if sisiA == sisiB && sisiB == sisiC {
	// 	fmt.Println("Segitiga sama sisi")
	// } else if sisiA == sisiB || sisiB == sisiC || sisiA == sisiC {
	// 	fmt.Println("Segitiga sama kaki")
	// } else {
	// 	fmt.Println("Segitiga sembarang")
	// }
	//make pseudocode
	//program Segitiga
	//kamus
	//sisiA, sisiB, sisiC : integer
	//algoritma
	//	input(sisiA, sisiB, sisiC)
	//	if sisiA == sisiB && sisiB == sisiC then
	//		output("Segitiga sama sisi")
	//	else if sisiA == sisiB || sisiB == sisiC || sisiA == sisiC then
	//		output("Segitiga sama kaki")
	//	else
	//		output("Segitiga sembarang")
	//	endif
	//endprogram

	//latihan4 Mutlak Absolut
	// var angka int
	// fmt.Scan(&angka)

	// if angka < 0 {
	//     angka = -angka
	// }
	// fmt.Println(angka)
	//make pseudocode
	//program Mutlak_Absolut
	//kamus
	//angka : integer
	//algoritma
	//	input(angka)
	//	if angka < 0 then
	//		angka = -angka
	//	endif
	//	output(angka)
	//endprogram

	//latihan5 Temperatur
	// var suhu1, suhu2, suhu3, suhu4, suhu5 float64
	// fmt.Scan(&suhu1, &suhu2, &suhu3, &suhu4, &suhu5)

	// if suhu2 > suhu1 && suhu3 > suhu2 && suhu4 > suhu3 && suhu5 > suhu4 {
	// 	fmt.Println("Stabil naik")
	// } else if suhu2 < suhu1 && suhu3 < suhu2 && suhu4 < suhu3 && suhu5 < suhu4 {
	// 	fmt.Println("Stabil turun")
	// } else {
	// 	fmt.Println("Tidak stabil")
	// }
	//make pseudocode
	//program Temperatur
	//kamus
	//suhu1, suhu2, suhu3, suhu4, suhu5 : real
	//algoritma
	//	input(suhu1, suhu2, suhu3, suhu4, suhu5)
	//	if suhu2 > suhu1 && suhu3 > suhu2 && suhu4 > suhu3 && suhu5 > suhu4 then
	//		output("Stabil naik")
	//	else if suhu2 < suhu1 && suhu3 < suhu2 && suhu4 < suhu3 && suhu5 < suhu4 then
	//		output("Stabil turun")
	//	else
	//		output("Tidak stabil")
	//	endif
	//endprogram

	//latihan6 Profit
	// var hasilBulanIni, hasilBulanLalu, profit float64
	// fmt.Scan(&hasilBulanIni, &hasilBulanLalu)

	// if hasilBulanIni > hasilBulanLalu {
	// 	profit = hasilBulanIni - hasilBulanLalu
	// 	fmt.Println("Penigkatan sebesar", profit)
	// } else if hasilBulanIni < hasilBulanLalu {
	// 	profit = hasilBulanLalu - hasilBulanIni
	// 	fmt.Println("Penurunan sebesar", profit)
	// } else {
	// 	fmt.Println("Tetap")
	// }
	//make pseudocode
	//program Profit
	//kamus
	//hasilBulanIni, hasilBulanLalu, profit real
	//algoritma
	//	input(hasilBulanIni, hasilBulanLalu)
	//	if hasilBulanIni > hasilBulanLalu then
	//		profit = hasilBulanIni - hasilBulanLalu
	//		output("Penigkatan sebesar", profit)
	//	else if hasilBulanIni < hasilBulanLalu then
	//		profit = hasilBulanLalu - hasilBulanIni
	//		output("Penurunan sebesar", profit)
	//	else
	//		output("Tetap")
	//	endif
	//endprogram

	//latihan7 LigaSepakBola
	// var gol1, gol2, gol3, gol4, goalTerbanyak, goalTerdikit int
	// fmt.Scan(&gol1, &gol2, &gol3, &gol4)

	// goalTerbanyak = gol1
	// goalTerdikit = gol1

	// if goalTerbanyak < gol2 {
	// 	goalTerbanyak = gol2
	// } else if gol2 < goalTerdikit {
	// 	goalTerdikit = gol2
	// }

	// if goalTerbanyak < gol3 {
	// 	goalTerbanyak = gol3
	// } else if gol3 < goalTerdikit {
	// 	goalTerdikit = gol3
	// }

	// if goalTerbanyak < gol4 {
	// 	goalTerbanyak = gol4
	// } else if gol4 < goalTerdikit {
	// 	goalTerdikit = gol4
	// }

	// fmt.Println(goalTerbanyak, goalTerdikit)
	//make pseudocode
	//program LigaSepakBola
	//kamus
	//gol1, gol2, gol3, gol4, goalTerbanyak, goalTerdikit : integer
	//algoritma
	//	input(gol1, gol2, gol3, gol4)
	//	goalTerbanyak = gol1
	//	goalTerdikit = gol1
	//	if goalTerbanyak < gol2 then
	//		goalTerbanyak = gol2
	//	else if gol2 < goalTerdikit then
	//		goalTerdikit = gol2
	//	endif
	//	if goalTerbanyak < gol3 then
	//		goalTerbanyak = gol3
	//	else if gol3 < goalTerdikit then
	//		goalTerdikit = gol3
	//	endif
	//	if goalTerbanyak < gol4 then
	//		goalTerbanyak = gol4
	//	else if gol4 < goalTerdikit then
	//		goalTerdikit = gol4
	//	endif
	//	output(goalTerbanyak, goalTerdikit)
	//endprogram

	//latihan8 Parkir
	// nanti dulu susah

	//latihan9 Akhir Tahun
	// var cashback, diskon, membership bool
	// var totalBelanja float64

	// fmt.Scan(&totalBelanja, &membership)
	// if totalBelanja >= 100.0 {
	// 	if membership == true {
	// 		diskon = true
	// 		if totalBelanja >= 200.0 {
	// 			cashback = true
	// 			totalBelanja = totalBelanja - (totalBelanja * 0.1) // 10%
	// 			totalBelanja = totalBelanja - 75000
	// 		} else {
	// 			cashback = false
	// 			totalBelanja = totalBelanja - (totalBelanja * 0.1)
	// 		}
	// 	} else {
	// 		diskon = true
	// 		totalBelanja = totalBelanja - (totalBelanja * 0.1)
	// 	}
	// }
	// fmt.Println("Kartu?", membership, "Diskon?", diskon, "Cashback?", cashback, "Total Belanja", totalBelanja)
	//make pseudocode
	//program AkhirTahun
	//kamus
	//cashback, diskon, membership : boolean
	//totalBelanja : real
	//algoritma
	//	input(totalBelanja, membership)
	//	if totalBelanja >= 100.0 then
	//		if membership == true then
	//			diskon = true
	//			if totalBelanja >= 200.0 then
	//				cashback = true
	//				totalBelanja = totalBelanja - (totalBelanja * 0.1) // 10%
	//				totalBelanja = totalBelanja - 75000
	//			else
	//				cashback = false
	//				totalBelanja = totalBelanja - (totalBelanja * 0.1)
	//			endif
	//		else
	//			diskon = true
	//			totalBelanja = totalBelanja - (totalBelanja * 0.1)
	//		endif
	//	endif
	//	output("Kartu?", membership, "Diskon?", diskon, "Cashback?", cashback, "Total Belanja", totalBelanja)
	//endprogram

	//latihan10 Avatar
	var bilP, appdewasaDiv, appdewasaMod, appkecilDiv, appkecilMod, takberangkat int
	fmt.Scan(&bilP)

	if bilP <= 15 {
		appdewasaDiv = bilP / 5
		appdewasaMod = bilP % 5

		if appdewasaMod != 0 {
			appdewasaDiv = appdewasaDiv + 1
		}
		fmt.Println("Dewasa", appdewasaDiv)
	} else if bilP > 15 {
		appdewasaDiv = 3

		if bilP <= 25 {
			appkecilDiv = (bilP - 15) / 2
			appkecilMod = (bilP - 15) % 2

			if appkecilMod != 0 {
				appkecilDiv = appkecilDiv + 1
			}
			fmt.Println("Dewasa",appdewasaDiv, "Kecil", appkecilDiv)
		} else if bilP > 25 {
			appkecilDiv = 5
			takberangkat = bilP - 25
			fmt.Println( "Dewasa", appdewasaDiv , "kecil", appkecilDiv, "tak berangkat", takberangkat)
		}
	}
	//make pseudocode
	//program Avatar
	//kamus
	//bilP, appdewasaDiv, appdewasaMod, appkecilDiv, appkecilMod, takberangkat : integer
	//algoritma
	//	input(bilP)
	//	if bilP <= 15 then
	//		appdewasaDiv = bilP / 5
	//      appdewasaMod = bilP % 5
	//		if appdewasaMod != 0 then
	//			appdewasaDiv = appdewasaDiv + 1
	//		endif
	//		output("Dewasa", appdewasaDiv)
	//	else if bilP > 15 then
	//		appdewasaDiv = 3
	//		if bilP <= 25 then
	//			appkecilDiv = (bilP - 15) / 2
	//			appkecilMod = (bilP - 15) % 2
	//			if appkecilMod != 0 then
	//				appkecilDiv = appkecilDiv + 1
	//			endif
	//			output("Dewasa",appdewasaDiv, "Kecil", appkecilDiv)
	//		else if bilP > 25 then
	//			appkecilDiv = 5
	//			takberangkat = bilP - 25
	//			output("Dewasa", appdewasaDiv , "kecil", appkecilDiv, "tak berangkat", takberangkat)
	//		endif
	//	endif
	//endprogram
}
