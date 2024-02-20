package main

import "fmt"

func main() {
	// this is a for loop code
	// jadi for loop itu, program akan berjalan selama kondisi yang ditentukan benar
	const n = 5
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}

	//this is a while loop code
	// jadi kondisi while itu selama kondisi awalnya true, maka akan berjalan, dan akan berhenti jika sudah
	// tidak memenuhi kondisi awal
	var tampung int = 2
	for tampung < 10 {
		fmt.Println(tampung)
		tampung++
	}
}
