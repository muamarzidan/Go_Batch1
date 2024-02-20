package main

import "fmt"

func constInt() {
	// jensi varibel
	const jenis1 int = 165 // this is constant variable that cannot be changed
	var jenis2 int // this is variable that can be changed

	// penulisan varibel dalam go
	var saya1 string
	saya1 = "John Lennon"

	var saya2 string = "John Lennon"

	saya3 := "John Lennon"


	fmt.Println(saya1, saya2, saya3, jenis1, jenis2) 
}


