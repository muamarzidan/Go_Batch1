package main

import "fmt"

func main() {
	var pilihLevel int
	var cekJawabanLevel1 = true
	// var cekJawabanLevel2 = true
	// var cekJawabanLevel3 = true

	var clueJawabanLevel1 = 3
	// var clueJawabanLevel2 = 2
	// var clueJawabanLevel3 = 1

	for pilihLevel < 1 || pilihLevel > 3 {
		fmt.Println("================================================================")
		fmt.Println("============ Selamat Datang di Tebak-Tebak Gramedia ============")
		fmt.Println("================================================================")
		fmt.Println("Pilih Level : ")
		fmt.Println("1. Level 1")
		fmt.Println("2. Level 3")
		fmt.Println("3. Level 5")
		fmt.Scan(&pilihLevel)
		fmt.Println("\t")

		if pilihLevel == 1 {
			for cekJawabanLevel1 {
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

				for Opsi < 1 || Opsi > 2 {
					fmt.Println("\nPilih Pertanyaan:")
					fmt.Println("1. Tebak jenis buku apa yang Roki baca")
					fmt.Println("2. Tebak negara terkecil di dunia")
					fmt.Println("\nPilihan Anda (1/2): ")
					fmt.Scan(&Opsi)

					if Opsi == 1 {
						for kesempatanJawab > 0 {
							fmt.Print("\n")
							fmt.Println("Pertanyaan : Tebak jenis buku apa yang Roki baca?")
							fmt.Println("Clue default : Ketika anda membaca buku yang berjenis ini, anda akan berpikir keras")
							fmt.Print("\n")
							fmt.Print("Masukan tebakan anda: ")
							fmt.Scan(&Jawaban)
		
							if Jawaban == "Filsafat" || Jawaban == "filsafat" || Jawaban == "FILSAFAT" {
								fmt.Println("**************************************************************************")
								fmt.Println("*     Jawaban anda benar, Jenis buku yang Roki baca adalah Filsafat      *")
								fmt.Println("*   Tukarkan jawaban anda untuk mendapatkan diskon saat pembelian buku   *")
								fmt.Println("**************************************************************************")
								cekJawabanLevel1 = false
								kesempatanJawab = 0
							} else {
								kesempatanJawab = kesempatanJawab - 1
								fmt.Println("Jawaban anda salah")
								if kesempatanJawab == 0 {
									fmt.Print("\n")
									fmt.Println("Sorry, kesempatan anda menjawab sudah habis")
									fmt.Println("Game Over")
									fmt.Print("\n")
									cekJawabanLevel1 = false
									Clue = 4
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
							fmt.Print("\n")
							fmt.Println("Pertanyaan : Tebak negara terkecil di dunia")
							fmt.Println("Clue default : Negara ini terletak di Eropa")
							fmt.Print("\n")
							fmt.Print("Masukan tebakan anda: ")
							fmt.Scan(&Jawaban)

							if Jawaban == "Vatican" || Jawaban == "vatican" || Jawaban == "VATICAN" {
								fmt.Println("**************************************************************************")
								fmt.Println("*       Jawaban anda benar, Negara terkecil di dunia adalah Vatican      *")
								fmt.Println("*   Tukarkan jawaban anda untuk mendapatkan diskon saat pembelian buku   *")
								fmt.Println("**************************************************************************")
								cekJawabanLevel1 = false
								kesempatanJawab = 0
							} else {
								kesempatanJawab = kesempatanJawab - 1
								fmt.Println("Jawaban anda salah")
								if kesempatanJawab == 0 {
									fmt.Print("\n")
									fmt.Println("Sorry, kesempatan anda menjawab sudah habis")
									fmt.Println("Game Over")
									fmt.Print("\n")
									cekJawabanLevel1 = false
									Clue = 4
								}
								if Clue < clueJawabanLevel1 {
									fmt.Println("Apakah anda ingin menggunakan gunakan clue? (y / t)")
									fmt.Scan(&pakaiClue)
									if pakaiClue == "y" || pakaiClue == "Y"{
										Clue = Clue + 1
										if Clue == 1 {
											fmt.Print("\n")
											fmt.Println("************************ Clue 1 *****************************")
											fmt.Println("*     Negara ini terletak di eropa dan di negara italia      *")
											fmt.Println("*************************************************************")
											fmt.Print("\n")
										} else if Clue == 2 {											
											fmt.Print("\n")
											fmt.Println("************************ Clue 2 *****************************")
											fmt.Println("*   Memiliki luas wilayah 0.44 km2 dan di dalam kota roma   *")
											fmt.Println("*************************************************************")
											fmt.Print("\n")
										} else if Clue == 3 {
											fmt.Print("\n")
											fmt.Println("************************ Clue 3 *****************************")
											fmt.Println("*           Negara ini dimulai dengan huruf V                *")
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
					}else {
						fmt.Println("Pilihan tidak valid. Silakan pilih 1 atau 2.")
					}
				}
			}
		} else {
			fmt.Println("Level Tidak Tersedia")
			fmt.Println("\t")
		}
    }
}



// func main() {
// 	var monster, shield int
// 	var HPBulba, ATKBulba int
// 	var HPSquirtle, ATKSquirtle int
// 	var HPPikachu, ATKPikachu int
// 	var HPGarga, ATKGarga int
// 	var decision, move string

// 	HPBulba = 100
// 	ATKBulba = 15
// 	HPSquirtle = 250
// 	ATKSquirtle = 30
// 	HPPikachu = 400
// 	ATKPikachu = 60
// 	HPGarga = 300
// 	ATKGarga = 40
// 	shield = 3

// 	fmt.Println("============== POKEMON BATTLE ARENA ==============")
// 	fmt.Println("Choose one of this monster to be your alliance: ")
// 	fmt.Println("1. Bulbasaur (HP: 100, ATK : 15)")
// 	fmt.Println("2. Squirtle  (HP: 250, ATK : 30)")
// 	fmt.Println("3. Pikachu   (HP: 400, ATK : 60)")
// 	fmt.Println("==================================================")
// 	fmt.Scanln(&monster)
// 	if monster == 1 {
// 		fmt.Println("============== POKEMON BATTLE ARENA ==============")
// 		fmt.Println("Alrightt, Lets Bulbasaur do it's magic")
// 		fmt.Println("Your opponent will be Garganacl, are you ready? (YES/NO)")
// 		fmt.Println("==================================================")
// 		fmt.Scanln(&decision)
// 		if decision == "YES" {
// 			fmt.Println("============== POKEMON BATTLE ARENA ==============")
// 			fmt.Println("============= BULBASAUR VS GARGANACL =============")
// 			fmt.Println("You can use Shield three times")
// 			for HPBulba > 0 && HPGarga > 0 {
// 				fmt.Println("Attack / Shield")
// 				fmt.Scanln(&move)
// 				if move == "Attack" {
// 					HPBulba = HPBulba - ATKGarga
// 					HPGarga = HPGarga - ATKBulba
// 					fmt.Println("Bulbasaur HP	:", HPBulba)
// 					fmt.Println("Garga HP	:", HPGarga)
// 					fmt.Println("")
// 				} else if move == "Shield" {
// 					if shield > 0 {
// 						shield = shield - 1
// 						fmt.Println("Shield remaining: ", shield)
// 						fmt.Println("Bulbasaur HP	:", HPBulba)
// 						fmt.Println("Garga HP	:", HPGarga)
// 						fmt.Println("")
// 					} else {
// 						fmt.Println("All Shield already used")
// 					}
// 				}
// 				fmt.Println("==================================================")
// 			}

// 			if HPBulba <= 0 {
// 				fmt.Println("		D E F E A T E D	")
// 				fmt.Println("==================================================")
// 			}
// 			if HPGarga <= 0 {
// 				fmt.Println("		V I C T O R Y")
// 				fmt.Println("==================================================")
// 			}
// 		} else {
// 			fmt.Println("SEE YOU NEXT GAME")
// 		}
// 	} else if monster == 2 {
// 		fmt.Println("============== POKEMON BATTLE ARENA ==============")
// 		fmt.Println("Alrightt, Lets Squirtle do it's magic")
// 		fmt.Println("Your opponent will be Garganacl, are you ready? (YES/NO)")
// 		fmt.Println("==================================================")
// 		fmt.Scanln(&decision)
// 		if decision == "YES" || decision == "yes" {
// 			fmt.Println("============== POKEMON BATTLE ARENA ==============")
// 			fmt.Println("============== SQUIRTLE VS GARGANACL =============")
// 			fmt.Println("You can use Shield three times")
// 			for HPSquirtle > 0 && HPGarga > 0 {
// 				fmt.Println("Attack / Shield")
// 				fmt.Scanln(&move)
// 				if move == "Attack" {
// 					HPSquirtle = HPSquirtle - ATKGarga
// 					HPGarga = HPGarga - ATKSquirtle
// 					fmt.Println("Squirtle HP	:", HPSquirtle)
// 					fmt.Println("Garga HP	:", HPGarga)
// 					fmt.Println("")
// 				} else if move == "Shield" {
// 					if shield > 0 {
// 						shield = shield - 1
// 						fmt.Println("Shield remaining: ", shield)
// 						fmt.Println("Squrirtle HP	:", HPSquirtle)
// 						fmt.Println("Garga HP	:", HPGarga)
// 						fmt.Println("")
// 					} else {
// 						fmt.Println("All Shield already used")
// 					}
// 				}
// 				fmt.Println("==================================================")
// 			}

// 			if HPSquirtle <= 0 {
// 				fmt.Println("		D E F E A T E D	")
// 				fmt.Println("==================================================")
// 			}
// 			if HPGarga <= 0 {
// 				fmt.Println("		V I C T O R Y")
// 				fmt.Println("==================================================")
// 			}
// 		} else {
// 			fmt.Println("SEE YOU NEXT GAME")
// 		}
// 	} else if monster == 3 {
// 		fmt.Println("============== POKEMON BATTLE ARENA ==============")
// 		fmt.Println("Alrightt, Lets Pikachu do it's magic")
// 		fmt.Println("Your opponent will be Garganacl, are you ready? (YES/NO)")
// 		fmt.Println("==================================================")
// 		fmt.Scanln(&decision)
// 		if decision == "YES" || decision == "yes" {
// 			fmt.Println("============== POKEMON BATTLE ARENA ==============")
// 			fmt.Println("============== PIKACHU VS GARGANACL ==============")
// 			fmt.Println("You can use Shield three times")
// 			for HPPikachu > 0 && HPGarga > 0 {
// 				fmt.Println("Attack / Shield")
// 				fmt.Scanln(&move)
// 				if move == "Attack" {
// 					HPPikachu = HPPikachu - ATKGarga
// 					HPGarga = HPGarga - ATKPikachu
// 					fmt.Println("Pikachu HP	:", HPPikachu)
// 					fmt.Println("Garga HP	:", HPGarga)
// 					fmt.Println("")
// 				} else if move == "Shield" {
// 					if shield > 0 {
// 						shield = shield - 1
// 						fmt.Println("Shield remaining: ", shield)
// 						fmt.Println("Pikachu HP	:", HPPikachu)
// 						fmt.Println("Garga HP	:", HPGarga)
// 						fmt.Println("")
// 					} else {
// 						fmt.Println("All Shield already used")
// 					}
// 				}
// 				fmt.Println("==================================================")
// 			}

// 			if HPPikachu <= 0 {
// 				fmt.Println("		D E F E A T E D	")
// 				fmt.Println("==================================================")
// 			}
// 			if HPGarga <= 0 {
// 				fmt.Println("		V I C T O R Y")
// 				fmt.Println("==================================================")
// 			}
// 		} else {
// 			fmt.Println("SEE YOU NEXT GAME")
// 		}
// 	}
// }
