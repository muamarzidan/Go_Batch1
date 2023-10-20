package main

import "fmt"

func main() {
	var nMahasiswa, jumlahWaktu, waktuHari int
	fmt.Scan(&nMahasiswa)

	for i := 0; i < nMahasiswa; i++ {
		fmt.Scan(&waktuHari)
		jumlahWaktu += waktuHari
	}

	rata_rata := float64(jumlahWaktu) / float64(nMahasiswa)
	fmt.Printf("%.3f\n", rata_rata)
}
