package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a) // scan is for normla input
	fmt.Scanf("%d", &a) // scanf is for input with format
	fmt.Scanln(&a) // scanln is for input with new line

	fmt.Println(a) // println is for print with new line
	fmt.Print(a)   // print is for print without new line
	fmt.Printf("%d", a) // printf is for print with format

	var tipeint int
	fmt.Scanf("%d", &tipeint)
	fmt.Printf("%d\n", tipeint) // %d is for int

	var tipefloat float64
	fmt.Scanf("%f", &tipefloat)
	fmt.Printf("%f\n", tipefloat) // %f is for float

	var tipeByte byte
	fmt.Scanf("%c", &tipeByte)
	fmt.Printf("%c\n", tipeByte) // %c is for char or byte or rune
}
