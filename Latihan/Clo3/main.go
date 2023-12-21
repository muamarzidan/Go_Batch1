package main

import "fmt"


func main() {
	//soal 1
	// var inputHeiji, inputSinci, selisih int;
	// fmt.Scan(&inputHeiji, &inputSinci)
	// selisih = inputSinci - inputHeiji
	// if selisih < 0 {
	// 	selisih = selisih * -1
	// }
	// fmt.Println("selisih angka heiji dan sinci adalah", selisih)
	//make pseudo code
	//program selisihabsolut
	//kamus
	//inputHeiji, inputSinci, selisih : integer
	//algoritma
	//input(inputHeiji, inputSinci)
	//selisih <- inputSinci - inputHeiji
	//if selisih < 0 then
	//	selisih <- selisih * -1
	//endif
	//output(selisih)
	//endprogram

	//soal 2
	// var inputHeiji, inputSinci, selisih int;

	// fmt.Scan(&inputHeiji, &inputSinci)
	// selisih = inputHeiji - inputSinci

	// if selisih%2 == 0 {
	// 	if selisih > 0 && selisih <= 45 {
	// 		fmt.Println("Sinci menang")
	// 	}
	// } else if selisih%2 == -1 {
	// 	if selisih < 0 && selisih >= -16 {
	// 		fmt.Println("Heiji menang")
	// 	}
	// } else if selisih == 0 {
	// 	fmt.Println("Permainan seri")
	// } else {
	// 	fmt.Println("sinci dan henji kalah")
	// }
	//make pseudo code
	//program selisihabsolut
	//kamus
	//inputHeiji, inputSinci, selisih : integer
	//algoritma
	//input(inputHeiji, inputSinci)
	//selisih <- inputHeiji - inputSinci
	//if selisih mod 2 = 0 then
	//	if selisih > 0 and selisih <= 45 then
	//		output("Sinci menang")
	//	endif
	//else if selisih mod 2 = -1 then
	//	if selisih < 0 and selisih >= -16 then
	//		output("Heiji menang")
	//	endif
	//else if selisih = 0 then
	//	output("Permainan seri")
	//else
	//	output("sinci dan henji kalah")
	//endif
	//endprogram

	//soal 3
	// var inputAB, inputBC, inputCD, inputAD int
	// fmt.Scan(&inputAB, &inputBC, &inputCD, &inputAD)

	// if inputAB <= 0 && inputBC <= 0 && inputCD <= 0 && inputAD <= 0 {
	// 	fmt.Println("Angka harus lebih dari 0")
	// } else {
	// 	if inputAB == inputBC && inputBC == inputCD && inputCD == inputAD {
	// 		fmt.Println("Angka dapat membentuk persegi")
	// 	} else if inputAB == inputCD && inputBC == inputAD {
	// 		fmt.Println("Angka dapat membentu persegi panjang")
	// 	} else {
	// 		fmt.Println("angka tidak dapat membentuk bidang apapun")
	// 	}
	// }
	//make pseudo code
	//program cekpersegi
	//kamus
	//inputAB, inputBC, inputCD, inputAD : integer
	//algoritma
	//input(inputAB, inputBC, inputCD, inputAD)
	//if inputAB <= 0 and inputBC <= 0 and inputCD <= 0 and inputAD <= 0 then
	//	output("Angka harus lebih dari 0")
	//else
	//	if inputAB = inputBC and inputBC = inputCD and inputCD = inputAD then
	//		output("Angka dapat membentuk persegi")
	//	else if inputAB = inputCD and inputBC = inputAD then
	//		output("Angka dapat membentu persegi panjang")
	//	else
	//		output("angka tidak dapat membentuk bidang apapun")
	//	endif
	//endif
	//endprogram

	// var hp, attack, defense, rumus4 float64
	// var star int
	// star = 0
	
	// fmt.Scan(&hp, &attack, &defense)
	// rumus4 = ((hp + attack + defense) / 45 ) * 100

	// if (hp >= 0 && attack >= 0 && defense >= 0) && (hp <= 15 && attack <= 15 && defense <= 15) {
	// 	if rumus4 >= 0 && rumus4 < 100 {
	// 		if rumus4 >= 0 && rumus4 <= 48.90 {
	// 			fmt.Println("Overall, your pokémon has room for improvement as far as battling goes.", star, "stars")
	// 		} else if rumus4 >= 51.10 && rumus4 <= 64.45 {
	// 			star = star + 1
	// 			fmt.Println("Overall, your pokémon is pretty decent!", star, "stars")
	// 		} else if rumus4 >= 66.60 && rumus4 <= 80 {
	// 			star = star + 2
	// 			fmt.Println("Overall, your pokémon is really strong!", star, "stars")
	// 		} else if rumus4 >= 82.2 && rumus4 <= 97.8 {
	// 			star = star + 3
	// 			fmt.Println("Overall, your pokémon looks like it can really battle with the best of them!", star, "stars")
	// 		} else if rumus4 == 100 {
	// 			star = star + 4
	// 			fmt.Println("Overall, your pokémon looks like it can really battle with the best of them!”", star, "stars")
	// 		} 
	// 	} else {
	// 		fmt.Println("not a pokemon")
	// 	}
	// } else {
	// 	fmt.Println("not a pokemon")
	// }
	//make pseudo code
	//program cekpokemon
	//kamus
	//hp, attack, defense, rumus4 : float
	//star : integer
	//algoritma
	//input(hp, attack, defense)
	//rumus4 <- ((hp + attack + defense) / 45 ) * 100
	//if (hp >= 0 and attack >= 0 and defense >= 0) and (hp <= 15 and attack <= 15 and defense <= 15) then
	//	if rumus4 >= 0 and rumus4 < 100 then
	//		if rumus4 >= 0 and rumus4 <= 48.90 then
	//			output("Overall, your pokémon has room for improvement as far as battling goes.", star, "stars")
	//		else if rumus4 >= 51.10 and rumus4 <= 64.45 then
	//			star <- star + 1
	//			output("Overall, your pokémon is pretty decent!", star, "stars")
	//		else if rumus4 >= 66.60 and rumus4 <= 80 then
	//			star <- star + 2
	//			output("Overall, your pokémon is really strong!", star, "stars")
	//		else if rumus4 >= 82.2 and rumus4 <= 97.8 then	
	//			star <- star + 3	
	//			output("Overall, your pokémon looks like it can really battle with the best of them!", star, "stars")
	//		else if rumus4 = 100 then
	//			star <- star + 4
	//			output("Overall, your pokémon looks like it can really battle with the best of them!”", star, "stars")
	//		endif
	//	else
	//		output("not a pokemon")
	//	endif
	//else
	//	output("not a pokemon")
	//endif
	//endprogram

	// diberikan sebuah lampu yang bisa nyala menggunakan 2 tongkat, sebuah tongkat mempunyai energinya masing-masing, nah lampu tidak akan menyala 
	// jika energi tongkat itu sama, contoh jika tongkat ke 1 ganjil dan tongkat ke 2 ganjil atau tongkat ke 1 genap dan tongkat ke 2 genap, maka tidak bisa menyala, jadi kedua tongkat harus 
	// berbeda nilai ganjil genapnya, tidak boleh ganjil ganjil atau genap genapconst
	// var tongkat1, tongkat2 int
	// fmt.Scan(&tongkat1, &tongkat2)

	// if (tongkat1%2 == 0 && tongkat2%2 == 0) || (tongkat1%2 == 1 && tongkat2%2 == 1) {
	// 	fmt.Println("tidak bisa menyala")
	// } else {
	// 	fmt.Println("bisa menyala")
	// }

	// diberikan sebuah studi kasus, untuk mengkomversi waktu jam menjadi bukan 24 jam, jadi hanya 12 saja waktunya, misal kan saya input 13:00 maka 
	// outputnya 1:00 "PM", jika 0  maka jam 12:00 "AM"
	// var jam, menit int
	// fmt.Scan(&jam, &menit)

	// if jam == 0 {
	// 	fmt.Println("12:", menit, "AM")
	// } else if jam == 12 {
	// 	fmt.Println("12:", menit, "PM")
	// } else if jam > 12 {
	// 	jam = jam - 12
	// 	fmt.Println(jam, ":", menit, "PM")
	// } else {
	// 	fmt.Println(jam, ":", menit, "AM")
	// }

	// diberikan studi kasus lagi, ada sebuah bengkel yang hanya buka dari jam 8 pagi sampai jam 8 malam atau 20.00, nah asumsikan semua yang masuk kebengkel itu dalam waktu operasi bengkel
	// tersebut, nah masukan terdiri dari jam masuk, menit masuk dan berapa servicenya, jika pengerjaan service ketika ditambahkan ke waktu menit dan menit ditambahkan ke jam melebih jam 8 malam/20.00, maka akan print "silahkan datang kemabali
	// lagi besok", jika service tidak sebaliknya, maka akan print waktu jam selesai dan menit selesai( dimana jam dan menit ini, sudah ditambahkan oleh waktu pengerjaan servicenya)

	// var jam, menit, service, menitbaru, menitsisa int
	// fmt.Scan(&jam, &menit, &service)

	// if service >= 0 && menit >= 0 && jam >= 0 {
	// 	if service >= 60{
	// 		menit = menit + service
	// 		if menit >= 60 {
	// 			menitbaru = menit / 60
	// 			menitsisa = menit - 60
	// 			jam = jam + menitbaru
	// 				if jam <= 20 {
	// 					fmt.Println(jam, ":", menitsisa)
	// 				} else {
	// 					fmt.Println("silahkan datang kemabali lagi besok")
	// 				}
	// 		} else {
	// 			jam = jam + menit
	// 			fmt.Println(jam, ":", menit)
	// 				if jam <= 20 {
	// 					fmt.Println(jam, ":", menit)
	// 				} else {
	// 					fmt.Println("silahkan datang kemabali lagi besok")
	// 				}
	// 			}
	// 	} else {
	// 		if menit >= 60 {
	// 			menitbaru = menit / 60
	// 			menitsisa = menit - 60
	// 			jam = jam + menitbaru
	// 				if jam <= 20 {
	// 					fmt.Println(jam, ":", menitsisa)
	// 				} else {
	// 					fmt.Println("silahkan datang kemabali lagi besok")
	// 				}
	// 		} else {
	// 			fmt.Println(jam, ":", menit)
	// 				if jam <= 20 {
	// 					fmt.Println(jam, ":", menit)
	// 				} else {
	// 					fmt.Println("silahkan datang kemabali lagi besok")
	// 				}
	// 			}
	// 	}
	// } else {
	// 		fmt.Println("inputan tidak valid")
	// }
}


