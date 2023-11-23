// program cekbesarkecil

// kamus 
// 	x, y : integer
// 	cek : boolean

// algoritma
// 	input(x)
// 	input(y)

// 	cek <- y > x
// 	output(y, cek)
// 	x <- y

// 	while x > 0 do 
// 		input(y)
// 		cek <- y > x
// 		output(y, cek)
// 		x <- y
// 	endwhile
// 	output("selesai")
// endprogram

//make a code about pseoudocode form code above

package main 

import "fmt"

func main () {
	var x, y int
	var cek bool

	fmt.Scan(&x)
	fmt.Scan(&y)

	cek = y > x
	fmt.Println(y, cek)
	x = y

	for x > 0 {
		fmt.Scan(&y)
		cek = y > x
		fmt.Println(y, cek)
		x = y
	}
	fmt.Println("selesai")
}