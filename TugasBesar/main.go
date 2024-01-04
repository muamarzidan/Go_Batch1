package main

import "fmt"

func main() {
	var program bool = true

	for program == true {
		var kodeGame bool
		var namaBuku, cekPunya, kode string
		var harga, total, uang, jumlahBuku, kembalian, opsiMenu, opsiKategori, opsiBuku, diskon, cashback int

		fmt.Println("\n================================================================")
		fmt.Println("============ Selamat Datang di Toko Buku Gramedia ==============")
		fmt.Println("================================================================")
		fmt.Print("\n")
		fmt.Println("Apa yang ingin anda lakukan? (1 / 2)")
		fmt.Println("1. Membeli buku")
		fmt.Println("2. Bermain Tebak-Tebakan")
		fmt.Print("\n")
		fmt.Scan(&opsiMenu)

		if opsiMenu == 1 || opsiMenu == 2{
			for program == true {
				if opsiMenu == 1 {
					fmt.Print("\n")
					fmt.Println("Berikut kategori buku yang kami jual: ")
					fmt.Println("1. Buku Filsafat")
					fmt.Println("2. Buku Agama")
					fmt.Println("3. Buku Coding")
					fmt.Println("4. Buku Mangga")
					fmt.Println("Masukan nomor kategori yang ingin anda beli: ")
					fmt.Scan(&opsiKategori)
					fmt.Print("\n")	

					for program == true {
						if opsiKategori > 0 && opsiKategori < 5 {
							if opsiKategori == 1 {
								fmt.Println("1. Dunia Shopie")
								fmt.Println("2. Filsafat Plato")
								fmt.Println("3. Sejarah Filsafat Barat")
								fmt.Println("4. Filosofi Teras")
								fmt.Println("5. Filsafat Aristoteles")
								fmt.Println("Masukan nomor buku yang ingin anda beli: ")
								fmt.Scan(&opsiBuku)
								fmt.Print("\n")
		
								for opsiBuku > 0 && opsiBuku < 6 {
									if opsiBuku == 1 {
										namaBuku = "Dunia Shopie"
										harga = 100000
									} else if opsiBuku == 2 {
										namaBuku = "Filsafat Plato"
										harga = 150000
									} else if opsiBuku == 3 {
										namaBuku = "Sejarah Filsafat Barat"
										harga = 200000
									} else if opsiBuku == 4 {
										namaBuku = "Filosofi Teras"
										harga = 250000
									} else if opsiBuku == 5 {
										namaBuku = "Filsafat Aristoteles"
										harga = 300000
									} else {
										fmt.Println("Pilihan tidak valid. Silakan pilih buku 1 - 5")
										opsiBuku = 0
									}
		
									if opsiBuku > 0 && opsiBuku < 6 {
										total = 0
										fmt.Println("Berapa buku yang ingin anda beli?")
										fmt.Scan(&jumlahBuku)
										fmt.Println("\t")
		
										total = harga * jumlahBuku
										fmt.Println("Total harga buku yang anda beli adalah : ", total)
			
										fmt.Println("Masukan uang anda: ")
										fmt.Scan(&uang)
										fmt.Print("\n")
			
										if uang < harga {
											for harga > uang {
												fmt.Println("Silahkan masukan uang dengan nominal lebih besar")
												fmt.Println("Masukan uang anda: ")
												fmt.Scan(&uang)
												fmt.Print("\n")
											}
											diskon = 10000
											cashback = 20000
		
											fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
											fmt.Scan(&cekPunya)
		
											if cekPunya == "y" || cekPunya == "Y" {
												kodeGame = true
		
												for kodeGame == true {
													fmt.Println("Silahkan masukan kode")
													fmt.Scan(&kode)
													if (kode == "filsafat" || kode == "Filsafat" || kode == "FILSAFAT" ) || (kode == "Vatican" || kode == "vatican" || kode == "VATICAN") {
														total = total - diskon
														kembalian = 0
														kembalian = uang - total
														fmt.Println("========================= STRUK ============================")
														fmt.Println("Anda telah membeli buku")
														fmt.Println("Buku :  ", namaBuku)
														fmt.Println("Harga : ", harga)
														fmt.Println("Jumlah Buku : ", jumlahBuku)
														fmt.Println("Anda mendapatkan diskon :", diskon)
														fmt.Println("Total Harga : ", total)
														fmt.Println("Kembalian : ", kembalian)
														fmt.Print("====================== Terima Kasih ========================\n")
		
														kodeGame = false
														program = false
													} else if (kode == "AtomicHabits" || kode == "Atomichabits" || kode == "atomichabits" || kode == "ATOMICHABITS") || (kode == "Islandia" || kode == "islandia" || kode == "ISLANDIA") {
														total = total - diskon
														kembalian = 0
														kembalian = uang - total
														kembalian = kembalian + cashback
														fmt.Println("========================= STRUK ============================")
														fmt.Println("Anda telah membeli buku")
														fmt.Println("Buku :  ", namaBuku)
														fmt.Println("Harga : ", harga)
														fmt.Println("Jumlah buku : ", jumlahBuku)
														fmt.Println("Anda mendapatkan diskon : ", diskon)
														fmt.Println("Anda mendapatkan cashback : ", cashback)
														fmt.Println("Total harga : ", total)
														fmt.Println("Kembalian : ", kembalian)
														fmt.Print("====================== Terima Kasih ========================\n")
		
														kodeGame = false
														program = false
													} else if (kode == "infiniti" || kode == "Infiniti" || kode == "INFINITI") || (kode == "Ichneumon" || kode == "ichneumon" || kode == "ICHNEUMON"){
														total = total - total
														kembalian = 0
														fmt.Println("========================= STRUK ============================")
														fmt.Println("Anda telah membeli buku")
														fmt.Println("Buku :  ", namaBuku)
														fmt.Println("Harga : ", harga)
														fmt.Println("Jumlah buku : ", jumlahBuku)
														fmt.Println("Anda mendapatkan buku secara gratis")
														fmt.Println("Total harga : ", total)
														fmt.Println("Uang anda kembali : ", uang)
														fmt.Print("====================== Terima Kasih ========================\n")
		
														kodeGame = false
														program = false
													} else {
														fmt.Println("Kode tidak valid, silakan masukan kode yang benar")
														kodeGame = true
													}
												}
											} else if cekPunya == "t" || cekPunya == "T" {
												kembalian = 0
												kembalian = uang - total
												fmt.Println("========================= STRUK ============================")
												fmt.Println("Anda telah membeli buku")
												fmt.Println("Buku :  ", namaBuku)
												fmt.Println("Harga : ", harga)
												fmt.Println("Jumlah Buku : ", jumlahBuku)
												fmt.Println("Total Harga : ", total)
												fmt.Println("Kembalian : ", kembalian)
												fmt.Print("====================== Terima Kasih ========================\n")
		
												opsiBuku = -1
												opsiKategori = -1
												opsiMenu = -1
												program = false
											} else {
												fmt.Println("Pilihan tidak valid, program berhenti")
		
												opsiBuku = -1
												opsiKategori = -1
												opsiMenu = -1
												program = false
											}
										} else if uang >= harga {
											diskon = 10000
											cashback = 20000
		
											fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
											fmt.Scan(&cekPunya)
		
											if cekPunya == "y" || cekPunya == "Y" {
												kodeGame = true
		
												for kodeGame == true {
													fmt.Println("Silahkan masukan kode")
													fmt.Scan(&kode)
													if (kode == "filsafat" || kode == "Filsafat" || kode == "FILSAFAT" ) || (kode == "Vatican" || kode == "vatican" || kode == "VATICAN") {
														total = total - 10000
														kembalian = 0
														kembalian = uang - total
														fmt.Println("========================= STRUK ============================")
														fmt.Println("Anda telah membeli buku")
														fmt.Println("Buku :  ", namaBuku)
														fmt.Println("Harga : ", harga)
														fmt.Println("Jumlah Buku : ", jumlahBuku)
														fmt.Println("Anda mendapatkan diskon :", diskon)
														fmt.Println("Total Harga : ", total)
														fmt.Println("Kembalian : ", kembalian)
														fmt.Print("====================== Terima Kasih ========================\n")
		
														kodeGame = false
														program = false
														opsiBuku = -1
													} else if (kode == "AtomicHabits" || kode == "Atomichabits" || kode == "atomichabits" || kode == "ATOMICHABITS") || (kode == "Islandia" || kode == "islandia" || kode == "ISLANDIA") {
														total = total - diskon
														kembalian = 0
														kembalian = uang - total
														kembalian = kembalian + cashback
														fmt.Println("========================= STRUK ============================")
														fmt.Println("Anda telah membeli buku")
														fmt.Println("Buku :  ", namaBuku)
														fmt.Println("Harga : ", harga)
														fmt.Println("Jumlah buku : ", jumlahBuku)
														fmt.Println("Anda mendapatkan diskon : ", diskon)
														fmt.Println("Anda mendapatkan cashback : ", cashback)
														fmt.Println("Total harga : ", total)
														fmt.Println("Kembalian : ", kembalian)
														fmt.Print("====================== Terima Kasih ========================\n")
		
														kodeGame = false
														program = false
													} else if (kode == "infiniti" || kode == "Infiniti" || kode == "INFINITI") || (kode == "Ichneumon" || kode == "ichneumon" || kode == "ICHNEUMON"){
														total = total - total
														kembalian = 0
														fmt.Println("========================= STRUK ============================")
														fmt.Println("Anda telah membeli buku")
														fmt.Println("Buku :  ", namaBuku)
														fmt.Println("Harga : ", harga)
														fmt.Println("Jumlah buku : ", jumlahBuku)
														fmt.Println("Anda mendapatkan buku secara gratis")
														fmt.Println("Total harga : ", total)
														fmt.Println("Uang anda kembali : ", uang)
														fmt.Print("====================== Terima Kasih ========================\n")
		
														kodeGame = false
														program = false
													} else {
														fmt.Println("Kode tidak valid, silakan masukan kode yang benar")
														kodeGame = true
													}
												}
											} else if cekPunya == "t" || cekPunya == "T" {
												kembalian = 0
												kembalian = uang - total
												fmt.Println("========================= STRUK ============================")
												fmt.Println("Anda telah membeli buku")
												fmt.Println("Buku :  ", namaBuku)
												fmt.Println("Harga : ", harga)
												fmt.Println("Jumlah Buku : ", jumlahBuku)
												fmt.Println("Total Harga : ", total)
												fmt.Println("Kembalian : ", kembalian)
												fmt.Print("====================== Terima Kasih ========================\n")
												
												program = false
												opsiBuku = -1
												opsiKategori = -1
												opsiMenu = -1
											} else {
												fmt.Println("Pilihan tidak valid, program berhenti")
		
												program = false
												opsiBuku = -1
												opsiKategori = -1
												opsiMenu = -1
											}
										} 
									} else {
										fmt.Println("Masukan pilihan buku dengan benar (1 - 5)")
										fmt.Print("\n")
									}
								}
								opsiBuku = 0
								fmt.Println("Masukan pilihan buku dengan benar (1 - 5)")
								fmt.Print("\n")
							}
						} else {
							fmt.Println("Masukan pilihan kategori dengan benar (1 - 4)")
							opsiMenu = 1
							program = false
						}
					}
				} 
			}
		}  else {
			fmt.Print("Masukan pilihan dengan benar (1 / 2)\n")
			program = true
		}
	}
}



// var input, awal, urutan, akhir int

// fmt.Println("Masukkan bilangan: ")
// fmt.Scan(&input)

// urutan = 0
// awal = input % 10
// input = input / 10

// for input > 0 {
// 	akhir = input % 10

// 	if awal < akhir {
// 		if urutan == 0 {
// 			urutan = 1
// 		} else if urutan == 2 {
// 			urutan = 0
// 			input = 0
// 		}
// 	} else if awal > akhir {
// 		if urutan == 0 {
// 			urutan = 2 
// 		} else if urutan == 1 {
// 			urutan = 0 
// 			input = 0
// 		}
// 	}

// 	awal = akhir
// 	input = input / 10
// }

// if urutan == 1 {
// 	fmt.Println("Bilangan tersebut ascending")
// } else if urutan == 2 {
// 	fmt.Println("Bilangan tersebut descending")
// } else {
// 	fmt.Println("Bilangan tersebut tidak urut")
// } 


