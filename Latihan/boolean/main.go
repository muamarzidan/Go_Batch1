package main

import "fmt"

func main() {
	// var k byte
	// var cekK bool
	// fmt.Scan(&k)

	// cekK = (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z') || (k >= '0' && k <= '9')
	// fmt.Println(cekK)

	// var tahun int
	// var cekTahun bool
	// fmt.Scanln(&tahun)

	// cekTahun = (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) 
	// fmt.Println(cekTahun)

	var sudutA, sudutB, sudutC int
	var cekSegitiga bool
	fmt.Scanln(&sudutA, &sudutB, &sudutC)

	cekSegitiga = sudutA+sudutB > sudutC && sudutA+sudutC > sudutB && sudutB+sudutC > sudutA
	fmt.Println(cekSegitiga)
}
