package main 

import "fmt"

func main () {
	var n, keliling, luas, panjang, lebar, i int

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&panjang, &lebar)
		luas = panjang * lebar
		keliling = 2 * (panjang + lebar)
		fmt.Println(luas, keliling)
	}
}

// program hitungbanyakpersegipanjang

// kamus
// 	n, keliling, luas, panjang, lebar : integer
// algoritma
// 	input(n)
// 	for i <- 0 to n do
// 		input(panjang, lebar)
// 		luas <- panjang * lebar
// 		keliling <- 2 * (panjang + lebar)
// 		output(luas, keliling)
// 	endfor
// endprogram