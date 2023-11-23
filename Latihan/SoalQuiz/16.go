// program cekstring
// kamus
// 	n, i : integer
// 	text : string
// algoritma
// 	input(n)
// 	input(text)
// 	for i <- 0 to n do
// 	output(text)
// 	endfor
// endprogram

package main

import "fmt"

func main () {
	var n, i int
	var text string
	fmt.Scan(&n)
	fmt.Scan(&text)
	for i = 0; i < n; i++ {
		fmt.Print(text)
	}
}