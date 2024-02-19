package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a) // scan is for normal input
	fmt.Scanf("%d", &a) // scanf is for input with format or type 
	fmt.Scanln(&a) // scanln is for input with new line

	fmt.Println(a) // println is for print with new line
	fmt.Print(a)   // print is for print without new line
	fmt.Printf("%d", a) // printf is for print with format or type
}
