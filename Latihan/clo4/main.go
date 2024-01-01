package main 

import "fmt"

func main () {
	var curahHujan string
    var Toha bool
	fmt.Scan(&curahHujan)

    if curahHujan == "sangat tinggi" {
        Toha = true
        if Toha == true {
            fmt.Println("macet")
        }
    } else if curahHujan == "tinggi" {
        fmt.Println("tidak macet")
    } else if curahHujan == "rendah" {
        fmt.Println("tidak macet")
    }

	// var hargaSuwir, hargaCakue, hargaAtiAmpela, hargaTelur, hargaPolos, totalHarga int
	// hargaSuwir = 3000
	// hargaCakue = 1500
	// hargaAtiAmpela = 4500
	// hargaTelur = 4000
	// hargaPolos = 6000	

	// var suwir, cakue, atiAmpela, telur bool
	// fmt.Scan(&suwir, &cakue, &atiAmpela, &telur)

	// totalHarga = hargaPolos 

	// if suwir {
	// 	totalHarga = totalHarga + hargaSuwir
	// }
	// if cakue {
	// 	totalHarga = totalHarga + hargaCakue
	// }
	// if atiAmpela {
	// 	totalHarga = totalHarga + hargaAtiAmpela
	// }
	// if telur {
	// 	totalHarga = totalHarga + hargaTelur
	// }

	// fmt.Println(totalHarga)

	// var warnaAwal, warnaDua string
    // fmt.Scan(&warnaAwal, &warnaDua)
    
    // if (warnaAwal == "merah" && warnaDua == "kuning") || (warnaAwal == "kuning" && warnaDua == "merah")  {
    //     fmt.Println("orange")
    // } else if (warnaAwal == "kuning" && warnaDua == "biru") || (warnaAwal == "biru" && warnaDua == "kuning")  {
    //     fmt.Println("hijau")
    // } else if (warnaAwal == "biru" && warnaDua == "merah") || (warnaAwal == "merah" && warnaDua == "biru")  {
    //     fmt.Println("ungu")
    // } else {
    //     fmt.Println("masukan warna sesuai inputan")
    // }

	// var input1, input2, input3, hasil, adikPintar int

	// fmt.Scan(&input1, &input2, &input3, &hasil)

	// adikPintar = input1 + input2 + input3

	// if adikPintar == hasil {
	// 	fmt.Println("benar")
	// } else {
	// 	fmt.Println("salah")
	// }

	var hujan, payung string
    fmt.Scan(&hujan, &payung)
    
    if hujan == "tidak" {
        fmt.Println("berangkat")
    } else if hujan == "ya" {
		if payung == "tidak" {
			fmt.Println("diam di rumah")
		} else if payung == "ya" {
			fmt.Println("berangkat")	
		} else {
		fmt.Println("masukan inputan sesuai soal")
		}
	} else {
		fmt.Println("masukan inputan sesuai soal")
	}
}