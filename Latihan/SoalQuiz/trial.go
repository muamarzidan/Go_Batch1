package main 

import "fmt"

func main () {
	var x, b1, b2, b3, b4, b5 int;

	fmt.Scan(&x);

	//x = 12345
	b1 = x / 10000; // x / 10000 = 1
	b2 = (x / 1000) % 10; // (x / 1000) = 12, 12 % 10 = 2
	b3 = (x / 100) % 10; // (x / 100) = 123, 123 % 10 = 3
	b4 = (x / 10) % 10; // (x / 10) = 1234, 1234 % 10 = 4
	b5 = x % 10;// x % 10 = 5
}


// func main () {
// 	var bil, i, hasil int
// 	hasil = 0
// 	fmt.Scan(&bil)
// 	for i = bil; i > 0; i-- {
// 		fmt.Print(i, "x")
// 		hasil = hasil + i
// 	}
// 	fmt.Println(hasil)
// }

// program urutan
// kamus
// 	bil, i, hasil : integer
// algoritma
// 	hasil <- 0
// 	input(bil)
// 	for i <- bil downto 0 do
// 		output(i, "x")
// 	endfor
// endprogram



//cek ganjil genap
// func main () {
// 	var bil int;
// 	var hasil bool;

// 	fmt.Scan(&bil);
// 	hasil = bil % 2 == 1;
// 	fmt.Println(hasil);

// 	hasil = bil % 2 == 0;
// 	fmt.Println(hasil);
// }

// program ganjilGenap
// kamus
// 	bil, hasil : integer
// algoritma
// 	input(bil)
// 	hasil <- bil mod 2 = 1
// 	output(hasil)

// 	hasil <- bil mod 2 = 0
// 	output(hasil)
// endprogram


//pisah 2 angka puluhan
// func main () {
// 	var bil, bilAwal, bilAkhir int;

// 	fmt.Scan(&bil);
// 	bilAwal = bil / 10;

// 	bilAkhir = bil % 10;
// 	fmt.Println(bilAkhir, bilAwal);
// }

// program pisahAngkaPuluhan
// kamus
// 	bil, bilAwal, bilAkhir : integer
// algoritma
// 	input(bil)
// 	bilAwal <- bil div 10
// 	bilAkhir <- bil mod 10
// 	output(bilAkhir, bilAwal)
// endprogram