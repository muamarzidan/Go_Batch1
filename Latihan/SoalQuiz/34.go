package main 

import "fmt"

func main () {
	var n, sisi, luas, kel int;

	fmt.Scan(&n);
	for i := 0; i < n; i++ {
		fmt.Scan(&sisi);
		luas = sisi * sisi;
		kel = sisi * 4;
		fmt.Println(luas, kel);
	}
}

// program tigadanempat
// kamus 
// 	n, sisi, luas, kel : integer
// algoritma
// 	input(n)
// 	for i <- 0 to n do
// 		input(sisi)
// 		luas <- sisi * sisi
// 		kel <- sisi * 4
// 		output(luas, kel)
// 	endfor
// endprogram
