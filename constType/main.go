package main

import "fmt"

func constInt() {
	// jensi varibel
	const MAXTINGGI int = 165
	var tinggi int 

	// penulisan varibel dalam go
	// var saya1 string
	// saya1 = "John Lennon"
	// var saya2 string = "John Lennon"
	// saya3 := "John Lennon"


	fmt.Println(MAXTINGGI - tinggi) // the result is 0, cause the value of var tinggi is 170, not 165 
}


