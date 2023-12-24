package main

import "fmt"

func main() {
	//Soal1 Faktor Persekutuan
	// var N, M, i int
	// fmt.Scan(&N, &M)

	// i = 1
	// for i <= N && i <= M {
	// 	if N%i == 0 && M%i == 0 {
	// 		fmt.Print(i, " ")
	// 	}
	// 	i++
	// }
	// fmt.Println()
	//make pseudocode
	//program Faktor_Persekutuan
	//kamus
	//	N, M, i : integer
	//algoritma
	//	input(N, M)
	//	i <- 1
	//	while i <= N && i <= M do
	//		if N mod i = 0 && M mod i = 0 then
	//			output(i)
	//		endif
	//		i <- i + 1
	//	endwhile
	//endprogram

	//Soal2
	// var n, m int
	// fmt.Scan(&n, &m)

	// for m != 0 {
	// 	n, m = m, n%m
	// }

	// fmt.Println(n)
	//make pseudocode
	//program Faktor_Persekutuan_Terbesar
	//kamus
	//	n, m : integer
	//algoritma
	//	input(n, m)
	//	while m != 0 do
	//		n, m <- m, n mod m
	//	endwhile
	//	output(n)
	//endprogram

	//Soal3 Persekutuan Terkecil 
	// var n, m int
	// fmt.Scan(&n, &m)

	// result := n
	// for result%m != 0 {
	// 	result += n
	// }

	// fmt.Println(result)
	//make pseudocode
	//program Kelipatan_Terkecil
	//kamus
	//	n, m, result : integer
	//algoritma
	//	input(n, m)
	//	result <- n
	//	while result mod m != 0 do
	//		result <- result + n
	//	endwhile
	//	output(result)
	//endprogram

	//Soal4 Museum 1
	// var n, total, count int
	// fmt.Scan(&n)

	// for n >= 0 && n <= 200 {
	// 	total += n
	// 	count++
	// 	fmt.Scan(&n)
	// }

	// if count == 0 {
	// 	fmt.Println(0, 0)
	// } else {
	// 	fmt.Println(count-1, float64(total)/float64(count))
	// }
	//make pseudocode
	//program Museum
	//kamus
	//	n, total, count : integer
	//algoritma
	//	input(n)
	//	while n >= 0 && n <= 200 do
	//		total <- total + n
	//		count <- count + 1
	//		input(n)
	//	endwhile
	//	if count = 0 then
	//		output(0, 0)
	//	else
	//		output(count-1, total/count)
	//	endif
	//endprogram

	//Soal5 Museum 2
	// var totalPengunjung, i int
	// var cekHari bool

	// totalPengunjung = 0

	// for i = 1; i <= 5; i++ {
	// 	cekHari = false

	// 	for !cekHari {
	// 		var pengunjung int
	// 		fmt.Printf("Hari %d : ", i)
	// 		fmt.Scan(&pengunjung)

	// 		if pengunjung > 0 && totalPengunjung+pengunjung <= 200 {
	// 			totalPengunjung += pengunjung
	// 			cekHari = true
	// 		}
	// 	}
	// }
	// fmt.Printf("Jumlah pengunjung: %d orang\n", totalPengunjung)
	//make pseudocode
	//program Museum
	//kamus
	//	totalPengunjung, i : integer
	//	cekHari : boolean
	//algoritma
	//	totalPengunjung <- 0
	//	for i <- 1 to 5 do
	//		cekHari <- false
	//		while cekHari = false do
	//			input(pengunjung)
	//			if pengunjung > 0 && totalPengunjung + pengunjung <= 200 then
	//				totalPengunjung <- totalPengunjung + pengunjung
	//				cekHari <- true
	//			endif
	//		endwhile
	//	endfor
	//	output(totalPengunjung)
	//endprogram

	//Soal6 Balap Mobil Mini
	// var x, Pemenang string;
	// var a, b, i int;

	// a, b = 0, 0
	// for i = 0; i < 10; i++ {
	// 	fmt.Scan(&x)
	// 	if x == "A" {
	// 		a++
	// 	} else if x == "B" {
	// 		b++
	// 	}

	// 	if a == 5 && Pemenang == "" {
	// 		Pemenang = "A"
	// 	} else if b == 5 && Pemenang == "" {
	// 		Pemenang = "B"
	// 	} else if b > 5 || a > 5 {
	// 		Pemenang = "tidak valid"
	// 	}
	// }

	// fmt.Println(Pemenang)
	//make pseudocode
	//program Balap_Mobil_Mini
	//kamus
	//	x, Pemenang : string
	//	a, b, i : integer
	//algoritma
	//	a, b <- 0
	//	for i <- 0 to 9 do
	//		input(x)
	//		if x = "A" then
	//			a <- a + 1
	//		else if x = "B" then
	//			b <- b + 1
	//		endif
	//		if a = 5 && Pemenang = "" then
	//			Pemenang <- "A"
	//		else if b = 5 && Pemenang = "" then
	//			Pemenang <- "B"
	//		else if b > 5 || a > 5 then
	//			Pemenang <- "tidak valid"
	//		endif
	//	endfor
	//	output(Pemenang)
	//endprogram

	//Soal7 Digit Terurut
	// var x, bilSeb, bilBaru int
	// var cekAscending, cekDescending bool

	// fmt.Scan(&x)

	// bilSeb = x % 10
	// x /= 10

	// cekAscending, cekDescending = true, true

	// for x > 0 {
	// 	bilBaru = x % 10
	// 	if bilSeb > bilBaru {
	// 		cekDescending = false
	// 	} else if bilSeb < bilBaru {
	// 		cekAscending = false
	// 	}

	// 	x /= 10
	// 	bilSeb = bilBaru
	// }

	// if cekAscending {
	// 	fmt.Println("ascending")
	// } else if cekDescending {
	// 	fmt.Println("descending")
	// } else {
	// 	fmt.Println("tidak terurut")
	// }
	//make pseudocode
	//program Digit_Terurut
	//kamus
	//	x, bilSeb, bilBaru : integer
	//	cekAscending, cekDescending : boolean
	//algoritma
	//	input(x)
	//	bilSeb <- x mod 10
	//	x <- x div 10
	//	cekAscending, cekDescending <- true, true
	//	while x > 0 do
	//		bilBaru <- x mod 10
	//		if bilSeb > bilBaru then
	//			cekDescending <- false
	//		else if bilSeb < bilBaru then
	//			cekAscending <- false
	//		endif
	//		x <- x div 10
	//		bilSeb <- bilBaru
	//	endwhile
	//	if cekAscending then
	//		output("ascending")
	//	else if cekDescending then
	//		output("descending")
	//	else
	//		output("tidak terurut")
	//	endif
	//endprogram

	//Soal8  MyTelU1
	// var x int;

	// fmt.Scan(&x)

	// if x == 0 {
	// 	x = 50
	// }
	
	// if x > 200 {
	// 	fmt.Println("Gold user")
	// } else if x >= 100 && x <= 200{
	// 	fmt.Println("Platinum user")
	// } else if x >= 50 && x < 100{
	// 	fmt.Println("Silver user")
	// }
	//make pseudocode
	//program MyTelU1
	//kamus
	//	x : integer
	//algoritma
	//	input(x)
	//	if x = 0 then
	//		x <- 50
	//	endif
	//	if x > 200 then
	//		output("Gold user")
	//	else if x >= 100 && x <= 200 then
	//		output("Platinum user")
	//	else if x >= 50 && x < 100 then
	//		output("Silver user")
	//	endif
	//endprogram

	//Soal9 MyTelU2
	// var x int;

	// for x < 50 {
	// 	fmt.Scan(&x)
	// }
	
	// if x > 200 {
	// 	fmt.Println("Gold user")
	// } else if x >= 100 && x <= 200{
	// 	fmt.Println("Platinum user")
	// } else if x >= 50 && x < 100{
	// 	fmt.Println("Silver user")
	// }
	//make pseudocode
	//program MyTelU2
	//kamus
	//	x : integer
	//algoritma
	//	while x < 50 do
	//		input(x)
	//	endwhile
	//	if x > 200 then
	//		output("Gold user")
	//	else if x >= 100 && x <= 200 then
	//		output("Platinum user")
	//	else if x >= 50 && x < 100 then
	//		output("Silver user")
	//	endif
	//endprogram

	//Soal10 MyTelU3
	// var x, n, i, gold, platinum, silver int

	// fmt.Scan(&n)

	// for i = 1; i <= n; i++ {
	// 	x = 0
	// 	for x < 50 {
	// 		fmt.Scan(&x)
	// 	}

	// 	if x > 200 {
	// 		gold++
	// 	} else if x >= 100 && x <= 200 {
	// 		platinum++
	// 	} else if x >= 50 && x < 100 {
	// 		silver++
	// 	}
	// }

	// fmt.Print("Gold user : ", gold, ", Platinum user : ", platinum, ", Silver user : ", silver)
	//make pseudocode
	//program MyTelU3
	//kamus
	//	x, n, i, gold, platinum, silver : integer
	//algoritma
	//	input(n)
	//	for i <- 1 to n do
	//		x <- 0
	//		while x < 50 do
	//			input(x)
	//		endwhile
	//		if x > 200 then
	//			gold <- gold + 1
	//		else if x >= 100 && x <= 200 then
	//			platinum <- platinum + 1
	//		else if x >= 50 && x < 100 then
	//			silver <- silver + 1
	//		endif
	//	endfor
	//emdprogram

	// Soal11 MyTelU4
	// var x, total, gold, platinum, silver int;

	// total = 0
	// for total < 500 {
	// 	x = 0
	// 	for x < 50 {
	// 		fmt.Scan(&x)
	// 	}

	// 	total = total + x

	// 	if x > 200 {
	// 		gold++
	// 	} else if x >= 100 && x <= 200{
	// 		platinum++
	// 	} else if x >= 50 && x < 100{
	// 		silver++
	// 	}
	// }

	// fmt.Print("Gold user : ", gold, ", Platinum user : ", platinum, ", Silver user : ", silver)
	//make pseudocode
	//program MyTelU4
	//kamus
	//	x, total, gold, platinum, silver : integer
	//algoritma
	//	total <- 0
	//	while total < 500 do
	//		x <- 0
	//		while x < 50 do
	//			input(x)
	//		endwhile
	//		total <- total + x
	//		if x > 200 then
	//			gold <- gold + 1
	//		else if x >= 100 && x <= 200 then
	//			platinum <- platinum + 1
	//		else if x >= 50 && x < 100 then
	//			silver <- silver + 1
	//		endif
	//	endwhile
	//endprogram

	//Soal12 MyTelU5
	var x, total, totGold, totPlatinum, totSilver, jumlahGold, jumlahPlatinum, jumlahSilver int
	var averageGold, averagePlatinum, averageSilver float64

	total = 0
	totGold = 0
	totPlatinum = 0
	totSilver = 0

	for total < 500 {
		x = 0
		for x < 50 {
			fmt.Scan(&x)
		}

		total = total + x

		if x > 200 {
			jumlahGold++
			totGold = totGold + x
		} else if x >= 100 && x <= 200 {
			jumlahPlatinum++
			totPlatinum = totPlatinum + x
		} else if x >= 50 && x < 100 {
			jumlahSilver++
			totSilver = totSilver + x
		}
	}

	averageGold = float64(totGold) / float64(jumlahGold)
	averagePlatinum = float64(totPlatinum) / float64(jumlahPlatinum)
	averageSilver = float64(totSilver) / float64(jumlahSilver)

	if averageGold != averageGold {
		averageGold = 0

		fmt.Print("Gold user : ", averageGold)
	} else {
		fmt.Printf("Gold user : %.2f", averageGold)
	}

	if averagePlatinum != averagePlatinum {
		averagePlatinum = 0

		fmt.Print(", Platinum user : ", averagePlatinum)
	} else {

		fmt.Printf(", Platinum user : %.2f", averagePlatinum)
	}

	if averageSilver != averageSilver {
		averageSilver = 0

		fmt.Print(", Silver user : ", averageSilver)
	} else {
		fmt.Printf(", Silver user : %.2f", averageSilver)
	}
	//make pseudocode
	//program MyTelU5
	//kamus
	//	x, total, totGold, totPlatinum, totSilver, jumlahGold, jumlahPlatinum, jumlahSilver : integer
	//	averageGold, averagePlatinum, averageSilver : float
	//algoritma
	//	total, totGold, totPlatinum, totSilver <- 0
	//	while total < 500 do
	//		x <- 0
	//		while x < 50 do
	//			input(x)
	//		endwhile
	//		total <- total + x
	//		if x > 200 then
	//			jumlahGold <- jumlahGold + 1
	//			totGold <- totGold + x
	//		else if x >= 100 && x <= 200 then
	//			jumlahPlatinum <- jumlahPlatinum + 1
	//			totPlatinum <- totPlatinum + x
	//		else if x >= 50 && x < 100 then
	//			jumlahSilver <- jumlahSilver + 1
	//			totSilver <- totSilver + x
	//		endif
	//	endwhile
	//	averageGold <- totGold / jumlahGold
	//	averagePlatinum <- totPlatinum / jumlahPlatinum
	//	averageSilver <- totSilver / jumlahSilver
	//	if averageGold != averageGold then
	//		averageGold <- 0
	//		output("Gold user : ", averageGold)
	//	else
	//		output("Gold user : ", averageGold)
	//	endif
	//	if averagePlatinum != averagePlatinum then
	//		averagePlatinum <- 0
	//		output(", Platinum user : ", averagePlatinum)
	//	else
	//		output(", Platinum user : ", averagePlatinum)
	//	endif
	//	if averageSilver != averageSilver then
	//		averageSilver <- 0
	//		output(", Silver user : ", averageSilver)
	//	else
	//		output(", Silver user : ", averageSilver)
	//	endif
	//endprogram
}