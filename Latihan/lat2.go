package main

import "fmt"

func main () {
	var nPersegi, nSisi int;
	fmt.Scan(&nPersegi, &nSisi);
	
	for i := 0; i < nPersegi; i++ {
	keliling := nSisi * 4;
	luas := nSisi * nSisi;

		fmt.Println("Keliling persegi: ", keliling);
		fmt.Println("Luas persegi: ", luas); 
	}
}