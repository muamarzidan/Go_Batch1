package main

import "fmt"

func main() {
	var program bool = true

	for program == true {
		var opsiMenu, opsiKategori, opsiBuku int

		fmt.Println("================================================================")
		fmt.Println("============ Selamat Datang di Toko Buku Gramedia ==============")
		fmt.Println("================================================================")
		fmt.Print("\n")
		fmt.Println("Apa yang ingin anda lakukan? (1 / 2)")
		fmt.Println("1. Membeli buku")
		fmt.Println("2. Bermain Tebak-Tebakan")
		fmt.Print("\n")

		for opsiMenu < 1 || opsiMenu > 2 {
			fmt.Scan(&opsiMenu)
			fmt.Println("\t")

			if opsiMenu == 1 {
				fmt.Print("\n")
				fmt.Println("Berikut kategori buku yang kami jual: ")
				fmt.Println("1. Buku Filsafat")
				fmt.Println("2. Buku Agama")
				fmt.Println("3. Buku Coding")
				fmt.Println("4. Buku Mangga")
				fmt.Print("\n")
				fmt.Println("Masukan nomor kategori yang ingin anda beli: ")

				for opsiKategori < 1 || opsiKategori > 4 {
					fmt.Scan(&opsiKategori)
					fmt.Print("\n")
					if opsiKategori == 1 {
						var namaBuku string
						var harga, total, uang, jumlahBuku, kembalian int
						fmt.Println("1. Dunia Shopie")
						fmt.Println("2. Filsafat Plato")
						fmt.Println("3. Sejarah Filsafat Barat")
						fmt.Println("4. Filosofi Teras")
						fmt.Println("5. Filsafat Aristoteles")
						fmt.Print("\n")
						fmt.Println("Masukan nomor buku yang ingin anda beli: ")

						for opsiBuku < 1 || opsiBuku > 5 {
							fmt.Scan(&opsiBuku)
							fmt.Print("\n")
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

							if opsiBuku != 0{
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
									var cekPunya, kode string
									var diskon, cashback int
									diskon = 10000
									cashback = 20000

									fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
									fmt.Scan(&cekPunya)

									if cekPunya == "y" || cekPunya == "Y" {
										var kodeGame bool
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

										program = false
									} else {
										fmt.Println("Pilihan tidak valid, program berhenti")
										program = false
									}
								} else if uang >= harga {
									var cekPunya, kode string
									var diskon, cashback int
									diskon = 10000
									cashback = 20000

									fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
									fmt.Scan(&cekPunya)

									if cekPunya == "y" || cekPunya == "Y" {
										var kodeGame bool
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
									} else {
										fmt.Println("Pilihan tidak valid, program berhenti")
										program = false
									}
								}
							}
						}
					} else if opsiKategori == 2 {
						var namaBuku string
						var harga, total, uang, jumlahBuku, kembalian int
						fmt.Println("1. Al-Quran")
						fmt.Println("2. Alkitab")
						fmt.Println("3. Tafsir Al-Quran")
						fmt.Println("4. Tafsir Alkitab")
						fmt.Println("5. 1000 Hadits Pilihan")

						fmt.Println("Masukan nomor buku yang ingin anda beli: ")

						for opsiBuku < 1 || opsiBuku > 5 {
							fmt.Scan(&opsiBuku)
							fmt.Print("\n")
							if opsiBuku == 1 {
								namaBuku = "Al-Quran"
								harga = 100000
							} else if opsiBuku == 2 {
								namaBuku = "Alkitab"
								harga = 150000
							} else if opsiBuku == 3 {
								namaBuku = "Hindunisme untuk Dummies"
								harga = 200000
							} else if opsiBuku == 4 {
								namaBuku = "Sirah Nabawiyah"
								harga = 250000
							} else if opsiBuku == 5 {
								namaBuku = "Paritta Suci"
								harga = 300000
							} else {
								fmt.Println("Pilihan tidak valid. Silakan pilih buku 1 - 5")
								opsiBuku = 0
							}

							if opsiBuku != 0{
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
									var cekPunya, kode string
									var diskon, cashback int
									diskon = 10000
									cashback = 20000

									fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
									fmt.Scan(&cekPunya)

									if cekPunya == "y" || cekPunya == "Y" {
										var kodeGame bool
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

										program = false
									} else {
										fmt.Println("Pilihan tidak valid, program berhenti")
										program = false
									}
								} else if uang >= harga {
									var cekPunya, kode string
									var diskon, cashback int
									diskon = 10000
									cashback = 20000

									fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
									fmt.Scan(&cekPunya)

									if cekPunya == "y" || cekPunya == "Y" {
										var kodeGame bool
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
									} else {
										fmt.Println("Pilihan tidak valid, program berhenti")
										program = false
									}
								}
							}
						}
					} else if opsiKategori == 3 {
						var namaBuku string
						var harga, total, uang, jumlahBuku, kembalian int
						fmt.Println("1. Go Lang")
						fmt.Println("2. Java")
						fmt.Println("3. C++")
						fmt.Println("4. Python")
						fmt.Println("5. Javascript")

						fmt.Println("Masukan nomor buku yang ingin anda beli: ")

						for opsiBuku < 1 || opsiBuku > 5 {
							fmt.Scan(&opsiBuku)
							fmt.Print("\n")
							if opsiBuku == 1 {
								namaBuku = "Go Lang"
								harga = 100000
							} else if opsiBuku == 2 {
								namaBuku = "Java"
								harga = 150000
							} else if opsiBuku == 3 {
								namaBuku = "C++"
								harga = 200000
							} else if opsiBuku == 4 {
								namaBuku = "Python"
								harga = 250000
							} else if opsiBuku == 5 {
								namaBuku = "Javascript"
								harga = 300000
							} else {
								fmt.Println("Pilihan tidak valid. Silakan pilih buku 1 - 5")
								opsiBuku = 0
							}

							if opsiBuku != 0{
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
									var cekPunya, kode string
									var diskon, cashback int
									diskon = 10000
									cashback = 20000

									fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
									fmt.Scan(&cekPunya)

									if cekPunya == "y" || cekPunya == "Y" {
										var kodeGame bool
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

										program = false
									} else {
										fmt.Println("Pilihan tidak valid, program berhenti")
										program = false
									}
								}
							}
						}
					} else if opsiKategori == 4 {
						var namaBuku string
						var harga, total, uang, jumlahBuku, kembalian int
						fmt.Println("1. One Piece")
						fmt.Println("2. Berserk")
						fmt.Println("3. Vagabond")
						fmt.Println("4. Slam Dunk")
						fmt.Println("5. Naruto")

						fmt.Println("Masukan nomor buku yang ingin anda beli: ")

						for opsiBuku < 1 || opsiBuku > 5 {
							fmt.Scan(&opsiBuku)
							fmt.Print("\n")
							if opsiBuku == 1 {
								namaBuku = "One Piece"
								harga = 100000
							} else if opsiBuku == 2 {
								namaBuku = "Berserk"
								harga = 150000
							} else if opsiBuku == 3 {
								namaBuku = "Vagabond"
								harga = 200000
							} else if opsiBuku == 4 {
								namaBuku = "Slam Dunk"
								harga = 250000
							} else if opsiBuku == 5 {
								namaBuku = "Naruto"
								harga = 300000
							} else {
								fmt.Println("Pilihan tidak valid. Silakan pilih buku 1 - 5")
								opsiBuku = 0
							}

							if opsiBuku != 0{
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
									var cekPunya, kode string
									var diskon, cashback int
									diskon = 10000
									cashback = 20000

									fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
									fmt.Scan(&cekPunya)

									if cekPunya == "y" || cekPunya == "Y" {
										var kodeGame bool
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

										program = false
									} else {
										fmt.Println("Pilihan tidak valid, program berhenti")
										program = false
									}
								} else if uang >= harga {
									var cekPunya, kode string
									var diskon, cashback int
									diskon = 10000
									cashback = 20000

									fmt.Println("Apakah anda mempunyai kode dari game tebak-tebakan? (y / t)")
									fmt.Scan(&cekPunya)

									if cekPunya == "y" || cekPunya == "Y" {
										var kodeGame bool
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
									} else {
										fmt.Println("Pilihan tidak valid, program berhenti")
										program = false
									}
								}
							}	
						}
					} else {
						fmt.Println("Pilihan tidak valid, program berhenti")
						program = false
					}
				}
			} else if opsiMenu == 2 {
				var pilihLevel int
				var cekJawabanLevel1 = true
				var cekJawabanLevel2 = true
				var cekJawabanLevel3 = true

				var clueJawabanLevel1 = 3
				var clueJawabanLevel2 = 2
				var clueJawabanLevel3 = 1

				for pilihLevel < 1 || pilihLevel > 3 {
					fmt.Println("\n================================================================")
					fmt.Println("============ Selamat Datang di Tebak-Tebak Gramedia ============")
					fmt.Println("================================================================")
					fmt.Println("Pilih Level : ")
					fmt.Println("1. Level 1")
					fmt.Println("2. Level 3")
					fmt.Println("3. Level 5")
					fmt.Scan(&pilihLevel)
					fmt.Println("\t")

					if pilihLevel == 1 {
						for cekJawabanLevel1 == true {
							fmt.Println("=================== Anda Memilih Level Pertama ==================")
							fmt.Println("=========================== Peraturan ===========================")
							fmt.Println("       1. Anda Mempunyai 5 Kali Kesempatan Untuk Menjawab        ")
							fmt.Println("     2. Anda Mempunyai 3 Kali Kesempatan Untuk Melihat Clue      ")
							fmt.Println("=================================================================")
							fmt.Println("=================================================================")

							var Opsi int
							var pakaiClue string
							var Jawaban string
							Clue := 0
							kesempatanJawab := 5

							for Opsi < 1 || Opsi > 3 {
								fmt.Println("\nPilih Pertanyaan:")
								fmt.Println("1. Tebak jenis buku apa yang Roki baca")
								fmt.Println("2. Tebak negara terkecil di dunia")
								fmt.Println("\nPilihan Anda (1 / 2): ")
								fmt.Scan(&Opsi)

								if Opsi == 1 {
									for kesempatanJawab > 0 {
										fmt.Println("\nPertanyaan : Tebak jenis buku apa yang Roki baca?")
										fmt.Println("Clue default : Ketika anda membaca buku yang berjenis ini, anda akan berpikir keras")
										fmt.Print("\nMasukan tebakan anda: ")
										fmt.Scan(&Jawaban)
					
										if Jawaban == "Filsafat" || Jawaban == "filsafat" || Jawaban == "FILSAFAT" {
											fmt.Println("**************************************************************************")
											fmt.Println("*     Jawaban anda benar, Jenis buku yang Roki baca adalah Filsafat      *")
											fmt.Println("*   Tukarkan jawaban anda untuk mendapatkan diskon saat pembelian buku   *")
											fmt.Println("**************************************************************************")
											cekJawabanLevel1 = false
											kesempatanJawab = 0
											program = false
										} else {
											kesempatanJawab = kesempatanJawab - 1
											fmt.Println("Jawaban anda salah")
											if kesempatanJawab == 0 {
												fmt.Println("\nSorry, kesempatan anda menjawab sudah habis")
												fmt.Println("Game Over")
												cekJawabanLevel1 = false
												Clue = 4
												kesempatanJawab = 0
												program = false
											}
											if Clue < clueJawabanLevel1 {
												fmt.Println("Apakah anda ingin menggunakan gunakan clue? (y / t)")
												fmt.Scan(&pakaiClue)
												if pakaiClue == "y" || pakaiClue == "Y"{
													Clue = Clue + 1
													if Clue == 1 {
														fmt.Print("\n")
														fmt.Println("************************ Clue 1 *****************************")
														fmt.Println("*   Jenis buku ini dibaca oleh orang berpengaruh di dunia   *")
														fmt.Println("*************************************************************")
														fmt.Print("\n")
													} else if Clue == 2 {											
														fmt.Print("\n")
														fmt.Println("************************ Clue 2 *****************************")
														fmt.Println("*              Jenis ini dibaca 3 kali ucapan               *")
														fmt.Println("*************************************************************")
														fmt.Print("\n")
													} else if Clue == 3 {
														fmt.Print("\n")
														fmt.Println("************************ Clue 3 *****************************")
														fmt.Println("*           Jenis buku ini dimulai dengan huruf F           *")
														fmt.Println("*************************************************************")
														fmt.Print("\n")
													}
												}
											}
											if kesempatanJawab > 0 {
												fmt.Println("Kesempatan menjawab anda tinggal", kesempatanJawab)
											}
										}
									}
								} else if Opsi == 2 {
									for kesempatanJawab > 0 {
										fmt.Println("\nPertanyaan : Tebak negara terkecil di dunia")
										fmt.Println("Clue default : Negara ini terletak di Eropa")
										fmt.Print("\nMasukan tebakan anda: ")
										fmt.Scan(&Jawaban)

										if Jawaban == "Vatican" || Jawaban == "vatican" || Jawaban == "VATICAN" {
											fmt.Println("**************************************************************************")
											fmt.Println("*       Jawaban anda benar, Negara terkecil di dunia adalah Vatican      *")
											fmt.Println("*   Tukarkan jawaban anda untuk mendapatkan diskon saat pembelian buku   *")
											fmt.Println("**************************************************************************")
											cekJawabanLevel1 = false
											kesempatanJawab = 0
											program = false
										} else {
											kesempatanJawab = kesempatanJawab - 1
											fmt.Println("Jawaban anda salah")
											if kesempatanJawab == 0 {
												fmt.Println("\nSorry, kesempatan anda menjawab sudah habis")
												fmt.Println("Game Over")
												cekJawabanLevel1 = false
												Clue = 4
												kesempatanJawab = 0
												program = false
											}
											if Clue < clueJawabanLevel1 {
												fmt.Println("Apakah anda ingin menggunakan gunakan clue? (y / t)")
												fmt.Scan(&pakaiClue)
												if pakaiClue == "y" || pakaiClue == "Y"{
													Clue = Clue + 1
													if Clue == 1 {
														fmt.Print("\n")
														fmt.Println("************************* Clue 1 *****************************")
														fmt.Println("*     Negara ini terletak di eropa dan di negara italia      *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													} else if Clue == 2 {											
														fmt.Print("\n")
														fmt.Println("************************* Clue 2 *****************************")
														fmt.Println("*   Memiliki luas wilayah 0.44 km2 dan di dalam kota roma    *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													} else if Clue == 3 {
														fmt.Print("\n")
														fmt.Println("************************* Clue 3 *****************************")
														fmt.Println("*           Negara ini dimulai dengan huruf V                *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													}
												}
											}
											if kesempatanJawab > 0 {
												fmt.Println("Kesempatan menjawab anda tinggal", kesempatanJawab)
											}
										}
									}
								} else {
									fmt.Println("\nPilihan tidak valid. Silakan pilih 1 atau 2.")
									Opsi = 0
								}
							}
						}
					} else if pilihLevel == 2 {
						for cekJawabanLevel2 == true {
							fmt.Println("=================== Anda Memilih Level Ketiga ==================")
							fmt.Println("=========================== Peraturan ==========================")
							fmt.Println("      1. Anda Mempunyai 4 Kali Kesempatan Untuk Menjawab       ")
							fmt.Println("    2. Anda Mempunyai 2 Kali Kesempatan Untuk Melihat Clue     ")
							fmt.Println("================================================================")
							fmt.Println("================================================================")

							var Opsi int
							var pakaiClue string
							var Jawaban string
							Clue := 0
							kesempatanJawab := 4

							for Opsi < 1 || Opsi > 2 {
								fmt.Println("\nPilih Pertanyaan:")
								fmt.Println("1. Tebak nama buku best seller di Gramedia")
								fmt.Println("2. Tebak negara")
								fmt.Println("\nPilihan Anda (1 / 2): ")
								fmt.Scan(&Opsi)

								if Opsi == 1 {
									for kesempatanJawab > 0 {
										fmt.Print("\n")
										fmt.Println("Pertanyaan : Tebak nama buku best seller di Gramedia")
										fmt.Println("Clue default : Buku ini ditulis oleh James Clear")
										fmt.Print("\n")
										fmt.Println("Note : Jawaban jangan dikasih spasi")
										fmt.Println("Masukan tebakan anda: ")
										fmt.Scan(&Jawaban)

										if Jawaban == "AtomicHabits" || Jawaban == "Atomichabits" || Jawaban == "atomichabits" || Jawaban == "ATOMICHABITS" {
											fmt.Println("**************************************************************************")
											fmt.Println("*         Jawaban anda benar, Buku tersebut bernama Atomic Habits        *")
											fmt.Println("*   Tukarkan jawaban anda untuk mendapatkan diskon saat pembelian buku   *")
											fmt.Println("**************************************************************************")
											cekJawabanLevel2 = false
											kesempatanJawab = 0
											program = false
										} else {
											kesempatanJawab = kesempatanJawab - 1
											fmt.Println("Jawaban anda salah")
											if kesempatanJawab == 0 {
												fmt.Println("\nSorry, kesempatan anda menjawab sudah habis")
												fmt.Println("Game Over")
												cekJawabanLevel2 = false
												Clue = 4
												kesempatanJawab = 0
												program = false
											}
											if Clue < clueJawabanLevel2 {
												fmt.Println("Apakah anda ingin menggunakan gunakan clue? (y / t)")
												fmt.Scan(&pakaiClue)
												if pakaiClue == "y" || pakaiClue == "Y"{
													Clue = Clue + 1
													if Clue == 1 {											
														fmt.Print("\n")
														fmt.Println("************************* Clue 1 *****************************")
														fmt.Println("*       Buku ini bisa membuat anda hidup lebih tertata       *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													} else if Clue == 2 {
														fmt.Print("\n")
														fmt.Println("************************* Clue 2 *****************************")
														fmt.Println("*     Buku ini dimulai dengan huruf A dan H dan 2 Kalimat    *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													}
												}
											}
											if kesempatanJawab > 0 {
												fmt.Println("Kesempatan menjawab anda tinggal", kesempatanJawab)
											}
										}
									}
								} else if Opsi == 2 {
									for kesempatanJawab > 0 {
										fmt.Print("\n")
										fmt.Println("Pertanyaan : Tebak negara")
										fmt.Println("Clue default : Negara ini tidak mempunyai nyamuk")
										fmt.Print("\n")
										fmt.Print("Masukan tebakan anda: ")
										fmt.Scan(&Jawaban)

										if Jawaban == "Islandia" || Jawaban == "islandia" || Jawaban == "ISLANDIA" {
											fmt.Println("**************************************************************************")
											fmt.Println("*              Jawaban anda benar, Negara tersebut Islandia              *")
											fmt.Println("*   Tukarkan jawaban anda untuk mendapatkan diskon saat pembelian buku   *")
											fmt.Println("**************************************************************************")
											cekJawabanLevel2 = false
											kesempatanJawab = 0
											program = false
										} else {
											kesempatanJawab = kesempatanJawab - 1
											fmt.Println("Jawaban anda salah")
											if kesempatanJawab == 0 {
												fmt.Println("\nSorry, kesempatan anda menjawab sudah habis")
												fmt.Println("Game Over")
												cekJawabanLevel2 = false
												Clue = 4
												kesempatanJawab = 0
												program = false
											}
											if Clue < clueJawabanLevel2 {
												fmt.Println("Apakah anda ingin menggunakan gunakan clue? (y / t)")
												fmt.Scan(&pakaiClue)
												if pakaiClue == "y" || pakaiClue == "Y"{
													Clue = Clue + 1
													if Clue == 1 {											
														fmt.Print("\n")
														fmt.Println("************************* Clue 1 *****************************")
														fmt.Println("*     Negara ini memiliki banyak gunung es dan gletser       *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													} else if Clue == 2 {
														fmt.Print("\n")
														fmt.Println("************************* Clue 2 *****************************")
														fmt.Println("*    Negara ini dimulai dengan huruf I dan 7 Kalimat         *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													}
												}
											}
											if kesempatanJawab > 0 {
												fmt.Println("Kesempatan menjawab anda tinggal", kesempatanJawab)
											}
										}
									}
								} else {
									fmt.Println("\nPilihan tidak valid. Silakan pilih 1 atau 2.")
									Opsi = 0
								}
							}
						}
					} else if pilihLevel == 3 {
						for cekJawabanLevel3 == true {
							fmt.Println("=================== Anda Memilih Level Kelima ==================")
							fmt.Println("=========================== Peraturan ==========================")
							fmt.Println("      1. Anda Mempunyai 3 Kali Kesempatan Untuk Menjawab       ")
							fmt.Println("    2. Anda Mempunyai 1 Kali Kesempatan Untuk Melihat Clue     ")
							fmt.Println("================================================================")
							fmt.Println("================================================================")

							var Opsi int
							var pakaiClue string
							var Jawaban string
							Clue := 0
							kesempatanJawab := 3

							for Opsi < 1 || Opsi > 2 {
								fmt.Println("\nPilih Pertanyaan:")
								fmt.Println("1. Tebak berapa buku yang ada di gramedia")
								fmt.Println("2. Tebak serangga")
								fmt.Println("\nPilihan Anda (1 / 2): ")
								fmt.Scan(&Opsi)

								if Opsi == 1 {
									for kesempatanJawab > 0 {
										fmt.Print("\n")
										fmt.Println("Pertanyaan : Tebak berapa buku yang ada di gramedia")
										fmt.Println("Clue default : Buku yang ada di gramedia lebih dari 1000")
										fmt.Print("\n")
										fmt.Println("Masukan tebakan anda: ")
										fmt.Scan(&Jawaban)

										if Jawaban == "infiniti" || Jawaban == "Infiniti" || Jawaban == "INFINITI" {
											fmt.Println("**************************************************************************")
											fmt.Println("*                Jawaban anda benar, jumlahnya infiniti                  *")
											fmt.Println("*   Tukarkan jawaban anda untuk mendapatkan diskon saat pembelian buku   *")
											fmt.Println("**************************************************************************")
											cekJawabanLevel3 = false
											kesempatanJawab = 0	
											program = false
										} else {
											kesempatanJawab = kesempatanJawab - 1
											fmt.Println("Jawaban anda salah")
											if kesempatanJawab == 0 {
												fmt.Println("\nSorry, kesempatan anda menjawab sudah habis")
												fmt.Println("Game Over")
												cekJawabanLevel3 = false
												Clue = 4
												kesempatanJawab = 0
												program = false
											}
											if Clue < clueJawabanLevel3 {
												fmt.Println("Apakah anda ingin menggunakan gunakan clue? (y / t)")
												fmt.Scan(&pakaiClue)
												if pakaiClue == "y" || pakaiClue == "Y"{
													Clue = Clue + 1
													if Clue == 1 {											
														fmt.Print("\n")
														fmt.Println("************************* Clue 1 *****************************")
														fmt.Println("*    Jawabannya memakai bahasa inggris, dan ular memutar     *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													}
												}
											}
											if kesempatanJawab > 0 {
												fmt.Println("Kesempatan menjawab anda tinggal", kesempatanJawab)
											}
										}
									}
								} else if Opsi == 2 {
									for kesempatanJawab > 0 {
										fmt.Print("\n")
										fmt.Println("Pertanyaan : Tebak serangga")
										fmt.Println("Clue default : Serangga ini bersifat parasit")
										fmt.Print("\n")
										fmt.Println("Masukan tebakan anda: ")
										fmt.Scan(&Jawaban)

										if Jawaban == "Ichneumon " || Jawaban == "ichneumon" || Jawaban == "ICHNEUMON" {
											fmt.Println("**************************************************************************")
											fmt.Println("*               Jawaban anda benar, serangga tersebut Ichneumon         *")
											fmt.Println("*   Tukarkan jawaban anda untuk mendapatkan diskon saat pembelian buku   *")
											fmt.Println("**************************************************************************")
											cekJawabanLevel3 = false
											kesempatanJawab = 0
											program = false
										} else {
											kesempatanJawab = kesempatanJawab - 1
											fmt.Println("Jawaban anda salah")
											if kesempatanJawab == 0 {
												fmt.Println("\nSorry, kesempatan anda menjawab sudah habis")
												fmt.Println("Game Over")
												cekJawabanLevel3 = false
												Clue = 4
												kesempatanJawab = 0
												program = false
											}
											if Clue < clueJawabanLevel3 {
												fmt.Println("Apakah anda ingin menggunakan gunakan clue? (y / t)")
												fmt.Scan(&pakaiClue)
												if pakaiClue == "y" || pakaiClue == "Y"{
													Clue = Clue + 1
													if Clue == 1 {											
														fmt.Print("\n")
														fmt.Println("************************* Clue 1 *****************************")
														fmt.Println("*    Serangga ini bersifat parasit dan berbentuk ular         *")
														fmt.Println("**************************************************************")
														fmt.Print("\n")
													}
												}
											}
											if kesempatanJawab > 0 {
												fmt.Println("Kesempatan menjawab anda tinggal", kesempatanJawab)
											}
										}
									}
								} else {
									fmt.Println("\nPilihan tidak valid. Silakan pilih 1 atau 2.")
									Opsi = 0
								}
							}
						}
					} else {
						fmt.Println("Pilihan tidak valid. Silakan pilih level 1 - 3")
						pilihLevel = 0
					}
				}	
			} else {
				fmt.Println("Pilihan tidak valid, ilakan pilih 1 atau 2")
				opsiMenu = 0
			}
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


