package main

import "fmt"

func main() {
	// noraml if with else
	var input int

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&input)

	if input > 10 {
		fmt.Println("Angka lebih dari 10")
	} else {
		fmt.Println("Angka kurang dari 10")
	}

	// if with else if
	if input > 10 {
		fmt.Println("Angka lebih dari 10")
	} else if input < 10 {
		fmt.Println("Angka kurang dari 10")
	} else {
		fmt.Println("Angka sama dengan 10")
	}

	// if with short statement
	if input := 10; input > 10 {
		fmt.Println("Angka lebih dari 10")
	} else {
		fmt.Println("Angka kurang dari 10")
	}

	// if with short statement and else if
	if input := 10; input > 10 {
		fmt.Println("Angka lebih dari 10")
	} else if input < 10 {
		fmt.Println("Angka kurang dari 10")
	} else {
		fmt.Println("Angka sama dengan 10")
	}

	// if with if in else
	if input > 10 {
		fmt.Println("Angka lebih dari 10")
	} else {
		if input < 10 {
			fmt.Println("Angka kurang dari 10")
		} else {
			fmt.Println("Angka sama dengan 10")
		}
	}
}
