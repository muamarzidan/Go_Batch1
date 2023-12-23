package main

import "fmt"

func main() {
	//soal1 Alphanumeric
	// var char byte
	// var cekAlphanumeric bool

	// cekAlphanumeric = true

	// fmt.Scanf("%c",&char)
	// cekAlphanumeric = char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9'

	// fmt.Println(cekAlphanumeric)
	//make pseudocode
	//program Alphanumeric
	//kamus
	// char : byte
	// cekAlphanumeric : bool
	//algoritma
	// cekAlphanumeric <- true
	// input char
	// cekAlphanumeric <- char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9'
	// output cekAlphanumeric
	//endprogram

	//soal2 Kabisat
	// var tahun int
	// var cekTahun bool
	// fmt.Scan(&tahun)

	// cekTahun = (tahun % 4 == 0 && tahun % 100 != 0) || (tahun % 400 == 0)
	// fmt.Print(cekTahun)
	//make pseudocode
	//program tahun_kabisat
	//kamus
	// tahun : int
	// cekTahun : bool
	//algoritma
	// input tahun
	// cekTahun <- (tahun % 4 == 0 && tahun % 100 != 0) || (tahun % 400 == 0)
	// output cekTahun
	//endprogram

	//soal3 Segitiga
	// var sudutA, sudutB, sudutC int
	// var cekSegitiga bool

	// fmt.Scanln(&sudutA, &sudutB, &sudutC)

	// cekSegitiga = sudutA+sudutB > sudutC && sudutA+sudutC > sudutB && sudutB+sudutC > sudutA
	// fmt.Println(cekSegitiga)
	//make pseudocode
	//program segitiga
	//kamus
	// sudutA, sudutB, sudutC : int
	// cekSegitiga : bool
	//algoritma
	// input sudutA, sudutB, sudutC
	// cekSegitiga <- sudutA+sudutB > sudutC && sudutA+sudutC > sudutB && sudutB+sudutC > sudutA
	// output cekSegitiga
	//endprogram

	//soal4 Mini Market
	// var total int
    // var inginKartu, dapatKartu, dapatDiskon, dapatCashback, Diskon, Cashback bool

    // fmt.Scan(&total)
    // fmt.Scan(&inginKartu)

    // Diskon = total >= 100000
    // Cashback = total >= 200000 && inginKartu

    // dapatKartu = inginKartu
    // dapatDiskon = Diskon
    // dapatCashback = Cashback

    // fmt.Println("Kartu?", dapatKartu)
    // fmt.Println("Diskon?", dapatDiskon)
    // fmt.Println("Cashback?", dapatCashback)
	//make pseudocode
	//program mini_market
	//kamus
	// total : int
	// inginKartu, dapatKartu, dapatDiskon, dapatCashback, Diskon, Cashback : bool
	//algoritma
	// input total, inginKartu
	// Diskon <- total >= 100000
	// Cashback <- total >= 200000 && inginKartu
	// dapatKartu <- inginKartu
	// dapatDiskon <- Diskon
	// dapatCashback <- Cashback
	// output dapatKartu, dapatDiskon, dapatCashback
	//endprogram

	//soal5 Mid Point

	//soal6 Dua Lingkaran
	// var a, b, tipus int
    // var cekIrisan bool

    // fmt.Scan(&a, &b, &tipus)

    // cekIrisan = a + b < tipus || a + tipus < b || tipus + b < a

    // fmt.Print(cekIrisan)
	//make pseudocode
	//program dua_lingkaran
	//kamus
	// a, b, tipus : int
	// cekIrisan : bool
	//algoritma
	// input a, b, tipus
	// cekIrisan <- a + b < tipus || a + tipus < b || tipus + b < a
	// output cekIrisan
	//endprogram

	//soal Pramuka
	// var n, i int;
    // var a, b, c, d, e, hasil bool;

	// fmt.Scan(&n)

	// hasil = true

    // for i = 0; i < n && hasil == true; i++ {
    //     fmt.Scan(&a, &b, &c, &d, &e)
    //     hasil = a == true && b == true && c == true && d == true && e == true
    // }
    // fmt.Println(hasil)
	//make pseudocode
	//program pramuka
	//kamus
	// n, i : int
	// a, b, c, d, e, hasil : bool
	//algoritma
	// input n
	// hasil <- true
	// for i <- 0 to n-1
	//     input a, b, c, d, e
	//     hasil <- a == true && b == true && c == true && d == true && e == true
	// output hasil
	//endprogram
}
